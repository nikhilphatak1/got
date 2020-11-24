package main

import (
    "fmt"
    "log"
    "path/filepath"
)

// Tree tree got object
type Tree struct {
    oid     []byte
    entries map[string]interface{}
}

// NewTree tree constructor
func NewTree() *Tree {
    tree := Tree{}
    tree.entries = make(map[string]interface{})
    return &tree
}

// BuildTree build a Merkle tree based on directory structure
func BuildTree(entries []*Entry) *Tree {
    root := NewTree()

    for _, singleEntry := range entries {
        root.AddEntry(singleEntry.ParentDirectories(), singleEntry)
    }

    return root
}

// AddEntry add entry to this tree using the list of parent directories
func (t *Tree) AddEntry(parents []string, entry *Entry) {
    if len(parents) == 0 {
        _, baseName := filepath.Split(entry.name)
        t.entries[baseName] = entry
    } else {
        _, baseName := filepath.Split(parents[0])
        var subTree *Tree
        var ok bool
        if t.entries[baseName] != nil {
            subTree, ok = t.entries[baseName].(*Tree)
            if !ok {
                log.Panicln("Unable to cast subtree to Tree")
            }
        } else {
            subTree = NewTree()
        }
        subTree.AddEntry(parents[1:], entry)
    }
}

// Type "tree"
func (t *Tree) Type() string {
    return "tree"
}

// ToString convert tree to string
func (t *Tree) ToString() string {
    resultString := ""
    for _, entry := range t.entries {
        tmpString := fmt.Sprintf("%s %s\x00%s", entry.Mode(), entry.name, entry.oid)
        resultString = resultString + tmpString
    }
    return resultString
}

// SetOid set blob oid
func (t *Tree) SetOid(oid []byte) {
    t.oid = oid
}

// GetOid set blob oid
func (t *Tree) GetOid() []byte {
    return t.oid
}
