package main

// StorableObject struct that can be written to gogit DB
type StorableObject interface {
	ToString() string
	SetOid(oid []byte)
	GetOid() []byte
	Type() string
}