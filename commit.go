package main

import (
    "encoding/hex"
    "fmt"
)

// Commit commit object conforms to StorableObject interface
type Commit struct {
    treeOid    []byte
    author  Author
    message string
    oid     []byte
}

// NewCommit Commit constructor
func NewCommit(treeOid []byte, author *Author, message string) *Commit {
    commit := Commit{}
    commit.treeOid = treeOid
    commit.author = *author
    commit.message = message
    return &commit
}

// Type returns "commit"
func (c *Commit) Type() string {
    return "commit"
}

// ToString convert to string
func (c *Commit) ToString() string {
    return fmt.Sprintf("tree %s\nauthor %s\ncommitter %s\n\n%s",
        hex.EncodeToString(c.treeOid), c.author.ToString(), c.author.ToString(), c.message)
}

// SetOid set commit oid
func (c *Commit) SetOid(oid []byte) {
    c.oid = oid
}

// GetOid commit oid
func (c *Commit) GetOid() []byte {
    return c.oid
}
