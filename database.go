package main

import (
    "fmt"
    "crypto/sha1"
    "path/filepath"
    "time"
    "math/rand"
)

// Database for storing info in .gogit/objects directory
type Database struct {
    pathname string
    letters []rune
}

// NewDatabase create and return a Database
func NewDatabase(pathname string) Database {
    database := Database{}
    database.pathname = pathname
    database.letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
    return database
}

// Store store the given blob
func (d Database) Store(object Blob) {
    // content: the data type (usually 'blob') then a space, then the number of data bytes
    // followed by a null byte \x00 in hex, then the data as a byte array encoded as a string
    // using the %s verb.
    content := fmt.Sprintf("%s %d\x00%s", object.Type(), len(object.data), object.data)

    h := sha1.New()
    h.Write(object.data)
    object.oid = h.Sum(nil)
    d.writeObject(object.oid, content)
}

func (d Database) writeObject(oid []byte, content string) {
    oidString := string(oid)
    targetPath := filepath.Join(d.pathname, oidString[0:2], oidString[2:])
    dirname := filepath.Join(d.pathname, oidString[0:2])
    tempPath := filepath.Join(dirname, "temp_object_" + d.tempName(6))
}

func (d Database) tempName(length int) string {
    randLetters := make([]rune, length)
    for i := range randLetters {
        randLetters[i] = d.letters[rand.Intn(len(d.letters))]
    }
    return string(randLetters)
}