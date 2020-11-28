package main

import (
    "bufio"
    "encoding/hex"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"
    "time"
)

func main() {
    argsWithoutProgName := os.Args[1:]
    if len(argsWithoutProgName) == 0 {
        log.Panicln("No got command given!")
    }

    //fmt.Println("Info: Running command", argsWithoutProgName[0])
    switch argsWithoutProgName[0] {
    case "init":
        gitInit(argsWithoutProgName[1:])
    case "commit":
        gitCommit(argsWithoutProgName[1:])
    default:
        log.Panicln(argsWithoutProgName[0], "is not a valid got command")
    }
}

func gitInit(argsWithoutInit []string) {
    // get the path either from command line args or using the cwd
    var rootPath string
    var err error
    if len(argsWithoutInit) == 0 {
        rootPath, err = os.Getwd()
        if err != nil {
            log.Panicln("Unable to get current working directory.", err)
        }
    } else {
        rootPath = argsWithoutInit[0]
        rootPath, err = filepath.Abs(rootPath)
        if err != nil {
            log.Panicln("Invalid path.", err)
        }
        var fileInfo os.FileInfo
        fileInfo, err = os.Stat(rootPath)
        if err != nil {
            log.Panicln("Unable to get current working directory.", err)
        }
        if !fileInfo.IsDir() {
            log.Panicln("Given path is not a directory")
        }
    }

    // make 'objects' and 'refs' directories
    gotPath := filepath.Join(rootPath, ".got")
    gotDirs := [2]string{"objects", "refs"}
    for _, p := range gotDirs {
        err = os.MkdirAll(filepath.Join(gotPath, p), 0777)
        if err != nil {
            log.Panicln("Unable to create directory for got metadata")
        }
    }

    fmt.Println("Info: Initializing in directory", rootPath) // for example /home/user
}

func gitCommit(argsWithoutCommit []string) {
    rootPath, err := os.Getwd()
    if err != nil {
        log.Panicln("Unable to get current working directory.", err)
    }
    gotPath := filepath.Join(rootPath, ".got")
    dbPath := filepath.Join(gotPath, "objects")
    workspace := NewWorkspace(rootPath)
    database := NewDatabase(dbPath)
    refs := NewRefs(gotPath)

    commitFilePaths := workspace.ListFilePaths()
    commitEntries := make([]Entry, len(commitFilePaths))
    for i, filename := range commitFilePaths {
        fileData := workspace.ReadFile(filename)
        blob := NewBlob(fileData)
        database.Store(blob)

        fileInfo := workspace.StatFile(filename)
        commitEntries[i] = *NewEntry(filename, hex.EncodeToString(blob.oid), fileInfo)
    }
    root := BuildTree(commitEntries)
    // TODO implement Tree.Traverse
    root.Traverse(() -> |tree| database.Store(tree)) // this doesn't make sense yet

    parent := refs.ReadHead()
    name := os.Getenv("GOT_AUTHOR_NAME")
    email := os.Getenv("GOT_AUTHOR_EMAIL")
    // TODO add check for if name or email of author is empty, print warning but don't exit
    author := NewAuthor(name, email, time.Now())
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter commit message: ")
    message, err := reader.ReadString('\n')
    if err != nil {
        log.Panicln("Unable to read commit message from Stdin.", err)
    }
    message = strings.TrimRight(message, "\r\n")

    commit := NewCommit(parent, tree.GetOid(), author, message)
    database.Store(commit)
    err = refs.UpdateHead(commit.GetOid())
    if (err != nil) {
        log.Panicln("Unable to update .got/HEAD.", err)
    }

    var rootPrefix string
    if parent == nil {
        rootPrefix = "(root-commit) "
    } else {
        rootPrefix = ""
    }

    ioutil.WriteFile(filepath.Join(gotPath, "HEAD"), []byte(hex.EncodeToString(commit.GetOid())), 0777)
    fmt.Printf("[%s%s] %s\n", rootPrefix, hex.EncodeToString(commit.GetOid()), message)
}
