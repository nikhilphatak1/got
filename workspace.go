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
        fi, err := os.Stat(path)
        if err != nil {
            fmt.Println(err)
            return nil
        }
        mode := fi.Mode();
        if !mode.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    if err != nil {
        panic(err)
    }
    // for _, filename := range files {
    //     fmt.Println("file", filename)
    // }
    return files
}

// ReadFile read the contents of the file in the file helper's directory
func (w Workspace) ReadFile(filepath string) []byte {
    //targetFilePath := filepath.Join(w.pathname, filename)
    fileContents, err := ioutil.ReadFile(filepath)
    if err != nil {
        fmt.Println("Error: Unable to read file.", err)
        panic(err)
    }
    return fileContents
}
