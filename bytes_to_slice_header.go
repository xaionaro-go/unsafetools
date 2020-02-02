package unsafetools

import (
	"reflect"
	"unsafe"
)

func BytesToSliceHeader(b []byte) *reflect.SliceHeader {
	return (*reflect.SliceHeader)(unsafe.Pointer(&b))
}
