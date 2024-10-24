package unsafetools

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Pointer is a constraint for pointers only.
type Pointer[T any] interface {
	*T
}

// FieldByName receives a pointer to a struct into `obj` and returns
// a pointer to the value of the field with name `fieldName` (even if it's a private field).
//
// It will panic if it was passed non-pointer or/and to non-structure to
// `obj`, or if field `fieldName` does not exist in the structure.
func FieldByName[F any, T any, PTR Pointer[T]](obj PTR, fieldName string) *F {
	return FieldByNameInValue(reflect.ValueOf(obj), fieldName).Interface().(*F)
}

// FieldByNameInValue does the same as FieldByName, but works with reflect.Value
// instead of interfaces.
func FieldByNameInValue(v reflect.Value, fieldName string) reflect.Value {
	elem := v.Elem()
	valuePointer := (unsafe.Pointer)(elem.FieldByName(fieldName).UnsafeAddr())

	structField, ok := elem.Type().FieldByName(fieldName)
	if !ok {
		panic(fmt.Sprintf(`unable to detect type of field %v`, fieldName))
	}

	return reflect.NewAt(structField.Type, valuePointer)
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
func FieldByPath[F any, T any, PTR Pointer[T]](obj PTR, path []string) *F {
	if len(path) == 0 {
		panic(`path is empty`)
	}

	return FieldByPathInValue(reflect.ValueOf(obj), path).Interface().(*F)
}

// FieldByPathInValue does the same as FieldByPath, but works with reflect.Value
// instead of interfaces.
func FieldByPathInValue(in reflect.Value, path []string) reflect.Value {
	value := in.Elem()
	var structField reflect.StructField
	for _, fieldName := range path {
		var ok bool
		structField, ok = value.Type().FieldByName(fieldName)
		if !ok {
			panic(fmt.Sprintf(`field %v does not exist: %v`, fieldName, path))
		}
		value = value.FieldByName(fieldName)
		if !value.IsValid() {
			panic(fmt.Sprintf(`field %v does not exist: %v`, fieldName, path))
		}
	}
	valuePointer := (unsafe.Pointer)(value.UnsafeAddr())

	return reflect.NewAt(structField.Type, valuePointer)
}
