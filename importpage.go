package main

import "fmt"

func GeneratePage(entry *Entry) string {
	s := `
	<html><head><meta name="go-import" content="%s %s %s"></head></html>
	`
	return fmt.Sprintf(s, entry.ImportPath, entry.RepoType, entry.RepoUrl)
}
