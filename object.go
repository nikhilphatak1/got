package main

// StorableObject struct that can be
type StorableObject interface {
	ToString() string
	BytesCount() int
	SetOid(oid string)
	Type() string
}