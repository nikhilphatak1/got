package main

// StorableObject struct that can be written to got DB
type StorableObject interface {
    ToString() string
    SetOid(oid []byte)
    GetOid() []byte
    Type() string
}
