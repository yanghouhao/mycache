package mycache

//ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte //b is read-only
}

//Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

//ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

//cloneBytes makes a slice and return the copy of the cache
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

func (v ByteView) String() string {
	return string(v.b)
}
