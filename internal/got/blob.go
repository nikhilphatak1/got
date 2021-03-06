package got

// Blob helper for blob manipulation
type Blob struct {
	data []byte // should this be []byte ?
	oid  []byte
}

// NewBlob create a new Blob struct
func NewBlob(data []byte) *Blob {
	blob := Blob{}
	blob.data = data
	return &blob
}

// Type returns "blob"
func (b *Blob) Type() string {
	return "blob"
}

// String convert to string
func (b *Blob) String() string {
	return string(b.data)
}

// SetOid set blob oid
func (b *Blob) SetOid(oid []byte) {
	b.oid = []byte(oid)
}

// GetOid blob oid
func (b *Blob) GetOid() []byte {
	return b.oid
}
