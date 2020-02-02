package unsafetools

import (
	"reflect"
	"unsafe"
)

func SetPointerToBytes(pointerPointer interface{}, b []byte) error {
	uintptrPointer, ok := pointerPointer.(*uintptr)
	if !ok {
		pointer := reflect.ValueOf(pointerPointer).Elem()
		requiredSize := pointer.Type().Elem().Size()
		if uintptr(len(b)) < requiredSize {
			return &ErrSliceTooShort{requiredSize, uintptr(len(b))}
		}
		uintptrPointer = (*uintptr)(unsafe.Pointer(pointer.UnsafeAddr()))
	}
	*uintptrPointer = uintptr(unsafe.Pointer(BytesToSliceHeader(b).Data))
	return nil
}
