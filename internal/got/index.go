package got

import "os"

// Index index of files
type Index struct {
	entries map[string]interface{}
	lockfile *Lockfile
	// TODO hash digest field here
}

// IndexEntry entry in the index
type IndexEntry struct {
	ctime, ctimeNsec, mtime, mtimeNsec, dev, ino, uid, gid, size int
	mode, oid, path string
}

const (
	regularMode = 0100644
	executableMode = 0100755
	maxPathLen = 0xfff
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
	// TODO mode = executableMode if stat is executable else regulatMode
	// TODO flags = min(len([]byte(targetPath)), maxPathLen)
	entry := IndexEntry{}
	entry.ctime = stat.ModTime().Second()
	entry.ctimeNsec = stat.ModTime().Nanosecond()
	entry.mtime = stat.ModTime().Second()
	entry.mtimeNsec = stat.ModTime().Nanosecond()
	// TODO other fields
	// also ctime is set to mtime currently
	// need to figure out how to get create time
	return &entry
}

// String convert IndexEntry
func (e *IndexEntry) String() string {
	// TODO
	return ""
}

// WriteUpdates write to index
func (i *Index) WriteUpdates() bool {
	lock := i.lockfile.HoldForUpdate()
	if !lock {
		return false
	}

	i.beginWrite()

	// TODO create header []byte containing, in order:
	// "DIRC" string (4 bytes)
	// "2" (padded to 8 bytes aka 32 bit)
	// len(i.entries)  (padded to 8 bytes aka 32 bit)
	// TODO i.write(header)

	for _, entry := range i.entries {
		i.write(entry.(*IndexEntry).String())
	}
	i.finishWrite()

	return true
}

func (i *Index) beginWrite() {
	// TODO set SHA digest here
}

func (i *Index) write(data string) {
	i.lockfile.Write(data)
	// TODO update i.digest with data
}

func (i *Index) finishWrite() {
	// TODO write digest to lockfile
	i.lockfile.Commit()
}