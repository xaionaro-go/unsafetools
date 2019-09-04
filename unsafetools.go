package unsafetools

import (
	"fmt"
	"reflect"
	"unsafe"
)

// FieldByName receives a pointer to a struct into `obj` and returns
// a pointer to the value of the field with name `fieldName` (even if it's a private field).
//
// It will panic if it was passed non-pointer or/and to non-structure to
// `obj`, or if field `fieldName` does not exist in the structure.
//
// Also you may pass a sample of an object of the same type as you expect to verify
// if it is correct (argument `sample`). If the type won't match it will panic. To avoid this
// check just pass a non-typed nil as `sample`.
func FieldByName(obj interface{}, fieldName string, sample interface{}) unsafe.Pointer {
	elem := reflect.ValueOf(obj).Elem()
	result := (unsafe.Pointer)(elem.FieldByName(fieldName).UnsafeAddr())
	if sample == nil {
		return result
	}

	sampleType := reflect.ValueOf(sample).Type()
	if sampleType.Kind() != reflect.Ptr {
		panic(fmt.Sprintf(`The provided verification sample is not a pointer, it is %T. While this function can return only pointers. So the verification has no sense.`, sample))
	}

	structField, ok := elem.Type().FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf(`unable to detect type of field %v of %T`, fieldName, obj))
	}
	if structField.Type != sampleType.Elem() {
		panic(fmt.Sprintf(`field %v of %T is *%T, not %T`, fieldName, obj, structField.Type.Name, sample))
	}
	return result
}
