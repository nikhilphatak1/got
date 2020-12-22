package got

import (
	"log"
	"os"
	"path/filepath"
)

// Entry got entry
type Entry struct {
	name string
	oid  string
	info os.FileInfo
	standardMode string
	executableMode string
}

// TODO extract constants to const blocks
// const (
//     standardMode =
// )

// NewEntry Entry constructor
func NewEntry(name string, oid string, info os.FileInfo) *Entry {
	entry := Entry{}
	entry.name = name
	entry.oid = oid
	entry.info = info
	entry.standardMode = "100644"
	entry.executableMode = "100755"
	return &entry
}

// ParentDirectories descending list of paths for this entry's name,
// starting from one dir name and approaching (but not including)
// the full entry name
func (e *Entry) ParentDirectories() []string {
	var dirList []string
	var parent string
	for parent != "/" {
		parent = filepath.Dir(e.name)
		dirList = append([]string{parent}, dirList...)
	}

	log.Println("List from ParentDirectories:")
	for _, j := range dirList {
		log.Println(j)
	}

	return dirList
}

// Mode entry mode
func (e *Entry) Mode() string {
	if e.info.Mode()&0100 != 0 {
		return e.executableMode
	} else {
		return e.standardMode
	}
}
