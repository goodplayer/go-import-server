package main

import (
	"bufio"
	"container/list"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Uri          string
	ImportPath   string
	RepoType     string
	RepoUrl      string
	RepoVendor   string
	RepoHomepage string
}

func (this *Entry) GetAction(c *gin.Context) {
	switch this.RepoVendor {
	case "github-http":
		c.String(http.StatusOK, GeneratePageGithubHttp(this))
	default:
		c.String(http.StatusInternalServerError, "repo vendor unknown.")
	}
}

func ReadEntries() ([]*Entry, error) {
	f, err := os.OpenFile("gitrepo.list", os.O_RDONLY, 0666)
	if os.IsNotExist(err) {
		return nil, err
	}

	l := list.New()

	r := bufio.NewReader(f)
	data, _, err := r.ReadLine()
	for {
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		en, e := processLine(string(data))
		if e != nil {
			return nil, e
		}
		if en != nil {
			l.PushBack(en)
		}
		data, _, err = r.ReadLine()
	}

	result := make([]*Entry, l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		result[i] = e.Value.(*Entry)
		i++
	}
	return result, nil
}

func processLine(line string) (*Entry, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}
	if line[0] == '#' {
		return nil, nil
	}
	elements := strings.Split(line, " ")
	entry := new(Entry)
	entry.Uri = elements[0]
	entry.ImportPath = elements[1]
	entry.RepoType = elements[2]
	entry.RepoUrl = elements[3]
	entry.RepoVendor = elements[4]
	entry.RepoHomepage = elements[5]
	if entry.RepoVendor != "github-http" { // currently only support "github-http"
		return nil, errors.New("repo vendor invalid. line=" + line)
	}
	return entry, nil
}
