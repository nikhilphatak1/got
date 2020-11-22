package main

import (
    "os"
)

// Entry got entry
type Entry struct {
    name string
    oid  string
    info os.FileInfo
    standardMode string
    executableMode string
}

// NewEntry Entry constructor
func NewEntry(name string, oid string, info os.FileInfo) *Entry {
    entry := Entry{}
    entry.name = name
    entry.oid = oid
    entry.info = info
    entry.standardMode = "100644"
    entry.executableMode = "100755"
    return &entry
}

// Mode entry mode
func (e *Entry) Mode() string {
    if e.info.Mode()&0100 != 0 {
        return e.executableMode
    } else {
        return e.standardMode
    }
}
