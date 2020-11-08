package main

// Entry gogit entry
type Entry struct {
	name string
	oid string
}

// NewEntry Entry constructor
func NewEntry(name string, oid string) *Entry {
	entry := Entry{}
	entry.name = name
	entry.oid = oid
	return &entry
}
