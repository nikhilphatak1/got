package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
)

// FileHelper helper for file manipulation
type FileHelper struct {
	pathname string
	ignoreList []string
}

// NewFileHelper create a new FileHelper
func NewFileHelper(pathname string) FileHelper {
	fileHelper := FileHelper{}
	fileHelper.ignoreList = []string{".", "..", ".gogit"}
	fileHelper.pathname = pathname
	return fileHelper
}

// ListFiles list files in the project
func ListFiles(f FileHelper) []string {
	var files []string
	err := filepath.Walk(f.pathname, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
	return files
}

// ReadFile read the contents of the file in the file helper's directory
func ReadFile(f FileHelper, path string) string {
	targetFilePath := filepath.Join(f.pathname, path)
	fileContents, err := ioutil.ReadFile(targetFilePath)
	if err != nil {
		fmt.Println("Error: Unable to read file.", err)
		panic(err)
	}
	return string(fileContents)
}