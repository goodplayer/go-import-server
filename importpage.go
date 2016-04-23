package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GeneratePageGithubHttp(entry *Entry) string {
	s := `
<!DOCTYPE html>
<html>
    <head>
        <meta name="go-import" content="%s %s %s">
        <meta name "go-source" content="%s %s %s/tree/master{/dir} %s/tree/master{/dir}/{file}#L{line}">
    </head>
</html>
	`
	return fmt.Sprintf(s, entry.ImportPath, entry.RepoType, entry.RepoUrl, entry.ImportPath, entry.RepoHomepage, entry.RepoHomepage, entry.RepoHomepage)
}

func RegisterAction(entry *Entry, r *gin.Engine) {
	r.GET(entry.Uri, entry.GetAction)
	r.GET(entry.Uri+"/*path", entry.GetAction)
}
