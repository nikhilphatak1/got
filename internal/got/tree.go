package got

import (
    "fmt"
    "log"
    "path/filepath"
    "reflect"
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
func (e *entrySorter) Swap(i, j int) {
    e.entries[i], e.entries[j] = e.entries[j], e.entries[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (e *entrySorter) Less(i, j int) bool {
    return e.by(&e.entries[i], &e.entries[j])
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

// Traverse recursively store all trees this tree contains
// and call the passed function f on each tree. As a rule,
// call f on the parent tree after calling it on the child tree.
func (t *Tree) Traverse(f func(*Tree)) {
    for _, potentialSubtree := range t.entries {
        if reflect.TypeOf(potentialSubtree).String() == "main.Tree" {
            potentialSubtree.(*Tree).Traverse(f)
        }
    }
    f(t)
}

// ToString convert tree to string
// TODO update this for new tree structure
func (t *Tree) ToString() string {
    resultString := ""
    for name, entry := range t.entries {
        var tmpString string
        if reflect.TypeOf(entry).String() == "main.*Tree" {
            tmpString = fmt.Sprintf(
                "%s %s\x00%s", entry.(*Tree).Mode(), name, entry.(*Tree).oid)
        } else if reflect.TypeOf(entry).String() == "main.*Entry" {
            tmpString = fmt.Sprintf(
                "%s %s\x00%s", entry.(*Entry).Mode(), name, entry.(*Entry).oid)
        }

        resultString = resultString + tmpString
    }
    return resultString
}

// Mode tree mode
func (t *Tree) Mode() string {
    return "40000" // directory mode
}

// SetOid set blob oid
func (t *Tree) SetOid(oid []byte) {
    t.oid = oid
}

// GetOid set blob oid
func (t *Tree) GetOid() []byte {
    return t.oid
}
