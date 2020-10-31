package main

import (
	"fmt"
    "os"
    "path/filepath"
)

func main() {
	fmt.Println("Go Git it.")
	argsWithoutProgName := os.Args[1:]
	if len(argsWithoutProgName) == 0 {
        fmt.Println("Error: No gogit command given!")
        os.Exit(1)
	}

	fmt.Println("Info: Running command", argsWithoutProgName[0])
	switch argsWithoutProgName[0] {
	case "init":
		gitInit(argsWithoutProgName[1:])
	default:
        fmt.Println("Error:", argsWithoutProgName[0], "is not a valid gogit command")
        os.Exit(1)
	}
}

func gitInit(argsWithoutInit []string) {
    // get the path either from command line args or using the cwd
    var path string
    var err error
	if len(argsWithoutInit) == 0 {
		path, err = os.Getwd()
		if err != nil {
            fmt.Println("Error: Unable to get current working directory.", err)
            os.Exit(1)
        }
	} else {
        path = argsWithoutInit[0]
        path, err = filepath.Abs(path)
        if err != nil {
            fmt.Println("Error: Invalid path.", err)
            os.Exit(1)
        }
        var fileInfo os.FileInfo
        fileInfo, err = os.Stat(path)
        if err != nil{
            fmt.Println("Error: Unable to get current working directory.", err)
            os.Exit(1)
        }
        if !fileInfo.IsDir() {
            fmt.Println("Error: Given path is not a directory")
            os.Exit(1)
        }
    }

    // make 'objects' and 'refs' directories
    gogitPath := filepath.Join(path, ".gogit/")
    gogitDirs := [2]string{"objects", "refs"}
    for _, p := range gogitDirs {
        err = os.MkdirAll(filepath.Join(gogitPath,p), 0777)
        if err != nil {
            fmt.Println("Error: Unable to create directory for gogit metadata")
            os.Exit(1)
        }
    }

    fmt.Println("Info: Initializing in directory", path) // for example /home/user
}
