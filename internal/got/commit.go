package got

import (
	"encoding/hex"
	"fmt"
)

// Commit commit object conforms to StorableObject interface
type Commit struct {
	parentOid []byte
	treeOid   []byte
	author    Author
	message   string
	oid       []byte
}

// NewCommit Commit constructor
func NewCommit(parent []byte, treeOid []byte, author *Author, message string) *Commit {
	commit := Commit{}
	commit.parentOid = parent
	commit.treeOid = treeOid
	commit.author = *author
	commit.message = message
	return &commit
}

// Type returns "commit"
func (c *Commit) Type() string {
	return "commit"
}

// String convert to string
func (c *Commit) String() string {
	var str string
	if c.parentOid == nil {
		str = fmt.Sprintf("tree %s\nauthor %s\ncommitter %s\n\n%s",
			hex.EncodeToString(c.treeOid), c.author.String(), c.author.String(), c.message)
	} else {
		str = fmt.Sprintf("tree %s\nparent %s\nauthor %s\ncommitter %s\n\n%s",
			hex.EncodeToString(c.treeOid), hex.EncodeToString(c.parentOid), c.author.String(),
			c.author.String(), c.message)
	}
	return str
}

// SetOid set commit oid
func (c *Commit) SetOid(oid []byte) {
	c.oid = oid
}

// GetOid commit oid
func (c *Commit) GetOid() []byte {
	return c.oid
}
