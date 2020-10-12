package main

import (
    "os"
    "fmt"
)

func main() {
    fmt.Println("Go Git it.")
    argsWithoutProgName := os.Args[1:]
    if len(argsWithoutProgName) == 0 {
        fmt.Println("Error: No gogit command given!")
    }
    path, err := os.Getwd()
    if err != nil {
        fmt.Println("Error: Unable to get current working directory: ", err)
    }
    fmt.Println("Info: Working in directory", path)  // for example /home/user
    fmt.Println("Info: Running command", argsWithoutProgName[0])
    switch argsWithoutProgName[0] {
    case "init": gitInit(argsWithoutProgName[1:])
    default: fmt.Println("Error:", argsWithoutProgName[0], "is not a valid gogit command")
    }
}

func gitInit(argsWithoutInit []string) {
    
}