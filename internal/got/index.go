package got

import "os"

// Index index of files
type Index struct {
    entries map[string]interface{}
    lockfile *Lockfile
}

// IndexEntry entry in the index
type IndexEntry struct {
    ctime, ctimeNsec, mtime, mtimeNsec, dev, ino, uid, gid, size int
    mode, oid, path string
}

const (
    regularMode = 0100644
    
)

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
    entry := CreateEntry(targetPath, oid, stat)
    i.entries[targetPath] = entry
}

// CreateEntry new entry for this index
func CreateEntry(targetPath string, oid string, stat os.FileInfo) *IndexEntry {
    // TODO
    mode = 
}