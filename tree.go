package main

import (
    "fmt"
    "log"
    "path/filepath"
    "sort"
)

// Tree tree got object
type Tree struct {
    oid     []byte
    entries map[string]interface{}
}

// By is the type of a "less" function that defines the ordering of its Entry arguments.
type By func(e1, e2 *Entry) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(entries []Entry) {
	es := &entrySorter{
		entries: entries,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(es)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type entrySorter struct {
	entries []Entry
	by      func(e1, e2 *Entry) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (e *entrySorter) Len() int {
	return len(e.entries)
}

// Swap is part of sort.Interface.
func (s *entrySorter) Swap(i, j int) {
	s.entries[i], s.entries[j] = s.entries[j], s.entries[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *entrySorter) Less(i, j int) bool {
	return s.by(&s.entries[i], &s.entries[j])
}

// NewTree tree constructor
func NewTree() *Tree {
    tree := Tree{}
    tree.entries = make(map[string]interface{})
    return &tree
}

// BuildTree build a Merkle tree based on directory structure
func BuildTree(entries []Entry) *Tree {
    name := func(e1, e2 *Entry) bool {
		return e1.name < e2.name
    }
    By(name).Sort(entries)
    root := NewTree()

    for _, singleEntry := range entries {
        root.AddEntry(singleEntry.ParentDirectories(), &singleEntry)
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

func (t *Tree) Traverse(f func (Tree)) {
    for _, baseName := t.entries {
        
    }
    f(*t)
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
