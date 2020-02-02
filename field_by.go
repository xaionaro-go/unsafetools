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
func FieldByName(obj interface{}, fieldName string) interface{} {
	elem := reflect.ValueOf(obj).Elem()
	valuePointer := (unsafe.Pointer)(elem.FieldByName(fieldName).UnsafeAddr())

	structField, ok := elem.Type().FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf(`unable to detect type of field %v of %T`, fieldName, obj))
	}

	return reflect.NewAt(structField.Type, valuePointer).Interface()
}

// FieldByPath receives a pointer to a struct into `obj` and returns
// a pointer to the value of the field by path `path` (even if there're private
// fields in the path).
//
// `path` is a series of field names which are used to recursively
// deep into the structure to get the required field.
//
// It will panic if it was passed non-pointer or/and to non-structure to
// `obj`, or if a field from the path does not exist in the structure.
func FieldByPath(obj interface{}, path []string) interface{} {
	if len(path) == 0 {
		panic(`path is empty`)
	}

	elem := reflect.ValueOf(obj).Elem()
	var structField reflect.StructField
	value := elem
	for _, fieldName := range path {
		var ok bool
		structField, ok = value.Type().FieldByName(fieldName)
		if !ok {
			panic(fmt.Sprintf(`field %v does not exist: %T %v`, fieldName, obj, path))
		}
		value = value.FieldByName(fieldName)
		if !value.IsValid() {
			panic(fmt.Sprintf(`field %v does not exist: %T %v`, fieldName, obj, path))
		}
	}
	valuePointer := (unsafe.Pointer)(value.UnsafeAddr())

	return reflect.NewAt(structField.Type, valuePointer).Interface()
}
