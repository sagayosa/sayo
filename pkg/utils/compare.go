package utils

import "reflect"

func CompareSlice(first interface{}, second interface{}) bool {
	if reflect.TypeOf(first) != reflect.TypeOf(second) {
		return false
	}

	return reflect.DeepEqual(first, second)
}
