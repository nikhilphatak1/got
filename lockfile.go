package main

import (
    "fmt"
    "os"
    "path/filepath"
)

// MissingParentError e
type MissingParentError struct {}
func (m MissingParentError) Error() string {
    return "Missing Parent"
}

// NoPermissionError e
type NoPermissionError struct {}
func (n NoPermissionError) Error() string {
    return "No Permission"
}

// StaleLockError e
type StaleLockError struct {}
func (s StaleLockError) Error() string {
    return "Stale Lock"
}


// Lockfile abstraction for managing access to files that should
// only be written to by one process at a time
type Lockfile struct {
    filePath string
    lockPath string
    lock *os.File
}

// NewLockfile Lockfile constructor
func NewLockfile(path string) *Lockfile {
    lockfile := Lockfile{}
    lockfile.filePath = path
    ext := filepath.Ext(path)
    // if no ext, filepath.Ext returns "" so len(ext) is 0
    // if ext exists, we truncate and apppend .lock
    lockfile.lockPath = path[:len(path)-len(ext)] + ".lock"
    lockfile.lock = nil
    return &lockfile
}

// HoldForUpdate attempt to acquire lock
func (l *Lockfile) HoldForUpdate() bool {
    if l.lock == nil {
        acquiredFile, err := os.Open(l.lockPath)
        if err != nil {
            return false
        }
        l.lock = acquiredFile
    }
    return true
}

// Write write
func (l *Lockfile) Write(content string) {
    err := l.raiseOnStaleLock()
    if err != nil {
        fmt.Println("Error: Stale lock.", err)
        panic(err)
    }
    _, err = l.lock.WriteString(content)
    if err != nil {
        fmt.Println("Error: Unable to write to lockfile.", err)
        panic(err)
    }
}

// Commit commit
func (l *Lockfile) Commit() {
    err := l.raiseOnStaleLock()
    if err != nil {
        fmt.Println("Error: Stale lock.", err)
        panic(err)
    }

    err = l.lock.Close()
    if err != nil {
        fmt.Println("Error: Unable to close lockfile.", err)
        panic(err)
    }

    err = os.Rename(l.lockPath, l.filePath)
    if err != nil {
        fmt.Println("Failed to rename lockfile.", err)
        panic(err)
    }
    l.lock = nil
}

func (l *Lockfile) raiseOnStaleLock() error {
    if l.lock == nil {
        return StaleLockError{}
    }
    return nil
}