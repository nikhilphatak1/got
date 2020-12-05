package got

import (
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Workspace file helper for file manipulation
type Workspace struct {
    pathname   string
    ignoreList []string
}

// NewWorkspace create a new FileHelper
func NewWorkspace(pathname string) Workspace {
    workspace := Workspace{}
    workspace.ignoreList = []string{".got", ".git"}
    workspace.pathname = pathname
    return workspace
}

// ListFilePaths list files in the project with relative paths from root
func (w *Workspace) ListFilePaths() []string {
    var files []string
    err := filepath.Walk(w.pathname, func(path string, info os.FileInfo, err error) error {
        fi, err := os.Stat(path)
        if err != nil {
            log.Println("Unable to stat", path, err)
            return nil
        }

        mode := fi.Mode()
        if !mode.IsDir() {
            for _, item := range w.ignoreList {
                if strings.Contains(path, item) {
                    return nil
                }
            }

            relativePath, err := filepath.Rel(w.pathname, path)
            if err != nil {
                log.Panicln("Unable to get relative path for", path)
            }

            files = append(files, relativePath)
        }
        return nil
    })
    if err != nil {
        panic(err)
    }
    return files
}

// ReadFile read the contents of the file in the file helper's directory
func (w *Workspace) ReadFile(targetFilePath string) []byte {
    if !filepath.IsAbs(targetFilePath) {
        targetFilePath = filepath.Join(w.pathname, targetFilePath)
    }

    fileContents, err := ioutil.ReadFile(targetFilePath)
    if err != nil {
        log.Panicln("Unable to read file.", err)
    }
    return fileContents
}

// StatFile syscall stat on file
func (w *Workspace) StatFile(targetFilePath string) os.FileInfo {
    if !filepath.IsAbs(targetFilePath) {
        targetFilePath = filepath.Join(w.pathname, targetFilePath)
    }

    fileInfo, err := os.Stat(targetFilePath)
    if err != nil {
        log.Panicln("Unable to stat file", targetFilePath, err)
    }
    return fileInfo
}
