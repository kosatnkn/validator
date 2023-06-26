package validator

import "reflect"

func isSlice(data interface{}) bool {
	value := reflect.ValueOf(data)

	// Check if the underlying type of the value is an array
	return value.Kind() == reflect.Array || value.Kind() == reflect.Slice
}
