package main

import (
	"fmt"
)

// Commit commit object conforms to StorableObject interface
type Commit struct {
	tree    Tree
	author  Author
	message string
	oid     []byte
}

// Type returns "commit"
func (c *Commit) Type() string {
	return "commit"
}

// ToString convert to string
func (c *Commit) ToString() string {
	return fmt.Sprintf("tree %s\nauthor %s\ncommitter %s\n\n%s",
		c.tree.ToString(), c.author.ToString(), c.author.ToString(), c.message)
}

// SetOid set blob oid
func (c *Commit) SetOid(oid []byte) {
	c.oid = []byte(oid)
}

// GetOid blob oid
func (c *Commit) GetOid() []byte {
	return c.oid
}
