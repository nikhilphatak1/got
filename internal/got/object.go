package got

// StorableObject struct that can be written to got DB
type StorableObject interface {
    String() string
    SetOid(oid []byte)
    GetOid() []byte
    Type() string
}
