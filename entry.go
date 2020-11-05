package main

// Entry gogit entry
type Entry struct {
	name string
	oid string
}

func NewEntry(name string, oid string) Entry {
	entry := Entry{}
	entry.name = name
	entry.oid = oid
	return entry
}
