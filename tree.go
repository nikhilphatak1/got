package main

import (
	"fmt"
)

// Tree tree gogit object
type Tree struct {
	oid string
	entries []Entry
	mode string
}

// NewTree tree constructor
func NewTree(entries []Entry) Tree {
	tree := Tree{}
	tree.entries = entries
	tree.mode = "100644"
	return tree
}

// Type "tree"
func (t Tree) Type() string {
	return "tree"
}

// ToString convert tree to string
func (t Tree) ToString() string {
	resultString := ""
	for _, entry := range t.entries {
		tmpString := fmt.Sprintf("%s %s\x00%s", t.mode, entry.name, entry.oid)
		resultString = resultString + tmpString
	}
	return resultString
}