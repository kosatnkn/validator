package validator

import (
	"reflect"
)

// isSlice checks whether the given unpacker is a slice or an array.
func isSlice(data interface{}) bool {
	value := reflect.ValueOf(data)

	// if data is a pointer, get the underlying value
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	// Check if the underlying type of the value is an array
	return value.Kind() == reflect.Array || value.Kind() == reflect.Slice
}
