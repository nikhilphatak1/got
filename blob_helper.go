package main

// BlobHelper helper for blob manipulation
type BlobHelper struct {
	data string // should be []byte ?
}

// NewBlobHelper create a new BlobHelper struct
func NewBlobHelper(data string) BlobHelper {
	blobHelper := BlobHelper{}
	blobHelper.data = data
	return blobHelper
}