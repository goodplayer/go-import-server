package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GeneratePage(entry *Entry) string {
	s := `
	<html><head><meta name="go-import" content="%s %s %s"></head></html>
	`
	return fmt.Sprintf(s, entry.ImportPath, entry.RepoType, entry.RepoUrl)
}

func RegisterAction(entry *Entry, r *gin.Engine) {
	r.GET(entry.Uri, entry.GetAction)
	r.GET(entry.Uri+"/*path", entry.GetAction)
}
