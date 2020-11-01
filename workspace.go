package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Workspace file helper for file manipulation
type Workspace struct {
	pathname   string
	ignoreList []string
}

// NewWorkspace create a new FileHelper
func NewWorkspace(pathname string) Workspace {
	workspace := Workspace{}
	workspace.ignoreList = []string{".", "..", ".gogit"}
	workspace.pathname = pathname
	return workspace
}

// ListFiles list files in the project
func (w Workspace) ListFiles() []string {
	var files []string
	err := filepath.Walk(w.pathname, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

// ReadFile read the contents of the file in the file helper's directory
func (w Workspace) ReadFile(filename string) string {
	targetFilePath := filepath.Join(w.pathname, filename)
	fileContents, err := ioutil.ReadFile(targetFilePath)
	if err != nil {
		fmt.Println("Error: Unable to read file.", err)
		panic(err)
	}
	return string(fileContents)
}
