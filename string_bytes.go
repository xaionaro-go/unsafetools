package unsafetools

import (
	"reflect"
	"unsafe"
)

// CastBytesToString casts a slice of bytes to a string in an unsafe way:
// the string is not guaranteed to be immutable (since it just reuses
// the same headers as the slice of bytes, which is modifiable).
//
// This is a zero-memory-allocation function.
//
//go:nosplit
func CastBytesToString(b []byte) string {
	return *(*string)((unsafe.Pointer)(&b))
}

// CastStringToBytes casts a string to a slice of bytes in an unsafe way:
// the string is not guaranteed to be immutable (since the returned slice
// has the same pointer underneath).
//
//go:nosplit
func CastStringToBytes(s string) []byte {
	hdr := (*reflect.StringHeader)((unsafe.Pointer)(&s))
	return *(*[]byte)((unsafe.Pointer)(&reflect.SliceHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
		Cap:  hdr.Len,
	}))
}
