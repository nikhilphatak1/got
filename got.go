package main

import (
    "bufio"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
)

func main() {
    //fmt.Println("Go Git it.")
    argsWithoutProgName := os.Args[1:]
    if len(argsWithoutProgName) == 0 {
        fmt.Println("Error: No got command given!")
        panic("No got command given!")
    }

    //fmt.Println("Info: Running command", argsWithoutProgName[0])
    switch argsWithoutProgName[0] {
    case "init":
        gitInit(argsWithoutProgName[1:])
    case "commit":
        gitCommit(argsWithoutProgName[1:])
    default:
        fmt.Println("Error:", argsWithoutProgName[0], "is not a valid got command")
        panic(fmt.Sprintf("Error: %s is not a valid got command", argsWithoutProgName[0]))
    }
}

func gitInit(argsWithoutInit []string) {
    // get the path either from command line args or using the cwd
    var rootPath string
    var err error
    if len(argsWithoutInit) == 0 {
        rootPath, err = os.Getwd()
        if err != nil {
            fmt.Println("Error: Unable to get current working directory.", err)
            panic(err)
        }
    } else {
        rootPath = argsWithoutInit[0]
        rootPath, err = filepath.Abs(rootPath)
        if err != nil {
            fmt.Println("Error: Invalid path.", err)
            panic(err)
        }
        var fileInfo os.FileInfo
        fileInfo, err = os.Stat(rootPath)
        if err != nil {
            fmt.Println("Error: Unable to get current working directory.", err)
            panic(err)
        }
        if !fileInfo.IsDir() {
            fmt.Println("Error: Given path is not a directory")
            panic("Error: Given path is not a directory")
        }
    }

    // make 'objects' and 'refs' directories
    gotPath := filepath.Join(rootPath, ".got")
    gotDirs := [2]string{"objects", "refs"}
    for _, p := range gotDirs {
        err = os.MkdirAll(filepath.Join(gotPath, p), 0777)
        if err != nil {
            fmt.Println("Error: Unable to create directory for got metadata")
            panic("Error: Unable to create directory for got metadata")
        }
    }

    fmt.Println("Info: Initializing in directory", rootPath) // for example /home/user
}

func gitCommit(argsWithoutCommit []string) {
    rootPath, err := os.Getwd()
    if err != nil {
        fmt.Println("Error: Unable to get current working directory.", err)
        panic(err)
    }
    gotPath := filepath.Join(rootPath, ".got")
    dbPath := filepath.Join(gotPath, "objects")
    workspace := NewWorkspace(rootPath)
    database := NewDatabase(dbPath)

    commitFilePaths := workspace.ListFilePaths()
    commitEntries := make([]*Entry, len(commitFilePaths))
    for i, filename := range commitFilePaths {
        fileData := workspace.ReadFile(filename)
        blob := NewBlob(fileData)
        database.Store(blob)
        commitEntries[i] = NewEntry(filename, hex.EncodeToString(blob.oid))
    }
    tree := NewTree(commitEntries)
    database.Store(tree)
    fmt.Println("stored tree in db")

    name := os.Getenv("got_AUTHOR_NAME")
    email := os.Getenv("got_AUTHOR_EMAIL")
    author := NewAuthor(name, email, time.Now())
    reader := bufio.NewReader(os.Stdin)
    message, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error: Unable to read commit message from Stdin.", err)
        panic(err)
    }
    fmt.Println("about to create commit")
    commit := NewCommit(tree.GetOid(), author, message)
    fmt.Println("about to store commit:", commit.ToString())
    database.Store(commit)

    fmt.Println("about to write HEAD")
    ioutil.WriteFile(filepath.Join(gotPath, "HEAD"), commit.GetOid(), 0777)
    fmt.Println("about to print root commit")
    fmt.Printf("(root-commit) %s %s\n", hex.EncodeToString(commit.GetOid()), message)
}
