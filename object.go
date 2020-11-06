package main

// StorableObject struct that can be written to gogit DB
type StorableObject interface {
	ToString() string
	BytesCount() int
	SetOid(oid []byte)
	GetOid() []byte
	Type() string
}