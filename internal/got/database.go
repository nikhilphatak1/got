package got

import (
	"bytes"
	"compress/gzip"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

// Database for storing info in .got/objects directory
type Database struct {
	pathname string
	letters  []rune
}

// NewDatabase create and return a Database
func NewDatabase(pathname string) Database {
	database := Database{}
	database.pathname = pathname
	database.letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	return database
}

// Store store the given blob
func (d Database) Store(object StorableObject) {
	objectString := object.String()
	// content: the data type (usually 'blob') then a space,
	// then the number of data bytes (number of runes) followed
	// by a null byte \x00 in hex, then the data as a byte array
	// encoded as a string using the %s verb.
	content := fmt.Sprintf("%s %d\x00%s", object.Type(), len(objectString), objectString)

	h := sha1.New()
	h.Write([]byte(content))
	object.SetOid(h.Sum(nil))
	d.writeObject(h.Sum(nil), content)
}

func (d Database) writeObject(oid []byte, content string) {
	rand.Seed(time.Now().UnixNano())
	oidString := hex.EncodeToString(oid)
	targetPath := filepath.Join(d.pathname, oidString[0:2], oidString[2:])
	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		// file exists, so do not overwrite
		return
	}

	dirname := filepath.Join(d.pathname, oidString[0:2])
	tempPath := filepath.Join(dirname, "temp_object_"+d.tempName(6))

	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		// dir dirname does not exist, so create it
		err = os.MkdirAll(dirname, 0777)
		if err != nil {
			log.Panicln("Unable to create directory for got metadata.", err)
		}
	}
	// compress the content string and write it to tempFile
	var b bytes.Buffer
	w, err := gzip.NewWriterLevel(&b, gzip.BestSpeed)
	if err != nil {
		log.Panicln("Invalid compression level.", err)
	}
	w.Write([]byte(content))
	w.Close()
	err = ioutil.WriteFile(tempPath, b.Bytes(), 0777)
	if err != nil {
		log.Panicln("Failed writing file.", err)
	}

	// rename tempFile path to targetPath and close the file descriptor (with error handling)
	err = os.Rename(tempPath, targetPath)
	if err != nil {
		log.Panicln("Failed to rename file.", err)
	}
}

func (d Database) tempName(length int) string {
	randLetters := make([]rune, length)
	for i := range randLetters {
		randLetters[i] = d.letters[rand.Intn(len(d.letters))]
	}
	return string(randLetters)
}
