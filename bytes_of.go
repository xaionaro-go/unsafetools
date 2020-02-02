package unsafetools

import (
	"reflect"
	"unsafe"
)

func BytesOf(obj interface{}) []byte {
	objValue := reflect.ValueOf(obj)
	objSize := objValue.Type().Elem().Size()
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: objValue.Elem().UnsafeAddr(),
		Len:  int(objSize),
		Cap:  int(objSize),
	}))
}
