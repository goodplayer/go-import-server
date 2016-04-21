package main

import (
	"bufio"
	"container/list"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Entry struct {
	Uri        string
	ImportPath string
	RepoType   string
	RepoUrl    string
}

func (this *Entry) GetAction(c *gin.Context) {
	c.String(http.StatusOK, GeneratePage(this))
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
		l.PushBack(en)
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
	elements := strings.Split(line, " ")
	entry := new(Entry)
	entry.Uri = elements[0]
	entry.ImportPath = elements[1]
	entry.RepoType = elements[2]
	entry.RepoUrl = elements[3]
	return entry, nil
}
