package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Refs abstraction for Got refs
type Refs struct {
	pathname string
}

// LockDeniedError e
type LockDeniedError struct {}
func (m LockDeniedError) Error() string {
    return "Lock Denied"
}

// NewRefs Refs constructor
func NewRefs(pathname string) *Refs {
	refs := Refs{}
	refs.pathname = pathname
	return &refs
}

// UpdateHead update .got/HEAD
func (r *Refs) UpdateHead(oid []byte) error {
	lockfile := NewLockfile(r.headPath())
	if !lockfile.HoldForUpdate() {
		return LockDeniedError{}
	}

	lockfile.Write(hex.EncodeToString(oid))
	lockfile.Write("\n")
	lockfile.Commit()
	return nil
}

// ReadHead return contents of .got/HEAD
func (r *Refs) ReadHead() []byte {
	if _, err := os.Stat(r.headPath()); !os.IsNotExist(err) {
		// HEAD file exists, so read and return
		contents, err := ioutil.ReadFile(r.headPath())
		if err != nil {
			fmt.Println("Error: Unable to read .got/HEAD file.", err)
			panic(err)
		}
		return contents
	}
	// HEAD does not exist, so return nil
	return nil
}

func (r *Refs) headPath() string {
	return filepath.Join(r.pathname, "HEAD")
}