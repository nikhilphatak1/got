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
		panic("No gogit command given!")
	}

	fmt.Println("Info: Running command", argsWithoutProgName[0])
	switch argsWithoutProgName[0] {
	case "init":
		gitInit(argsWithoutProgName[1:])
	case "commit":
		gitCommit(argsWithoutProgName[1:])
	default:
		fmt.Println("Error:", argsWithoutProgName[0], "is not a valid gogit command")
		panic(fmt.Sprintf("Error: %s is not a valid gogit command", argsWithoutProgName[0]))
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
	gogitPath := filepath.Join(rootPath, ".gogit")
	gogitDirs := [2]string{"objects", "refs"}
	for _, p := range gogitDirs {
		err = os.MkdirAll(filepath.Join(gogitPath, p), 0777)
		if err != nil {
			fmt.Println("Error: Unable to create directory for gogit metadata")
			panic("Error: Unable to create directory for gogit metadata")
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
	gogitPath := filepath.Join(rootPath, ".gogit")
	dbPath := filepath.Join(gogitPath, "objects")
	workspace := NewWorkspace(rootPath)
	database := NewDatabase(dbPath)

	for _, filename := range workspace.ListFiles() {
		fileData := workspace.ReadFile(filename)
        blob := NewBlob(fileData)
        database.Store(blob)
	}
}
