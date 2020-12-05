package got

import "os"

// Index index of files
type Index struct {
    entries map[string]interface{}
    lockfile *Lockfile
}

// NewIndex index constructor
func NewIndex(indexPath string) *Index {
    index := Index{}
    index.entries = make(map[string]interface{})
    index.lockfile = NewLockfile(indexPath)
    return &index
}

// Add add to the index
func (i *Index) Add(targetPath string, oid string, stat os.FileInfo) {
    // TODO make this call index's CreateEntry rather than NewEntry
    entry := NewEntry(targetPath, oid, stat)
    i.entries[targetPath] = entry
}

// CreateEntry new entry for this index
func CreateEntry(targetPath string, oid string, stat os.FileInfo) {
    // TODO
}