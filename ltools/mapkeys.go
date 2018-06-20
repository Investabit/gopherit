package ltools

import (
	"fmt"
	"reflect"
)

// MapKeys reports a slice of map keys as strings.
func MapKeys(mapToGetKeysFrom interface{}) (keys []string) {
	keys = []string{}
	mapType := reflect.TypeOf(mapToGetKeysFrom)
	mapKeyType := mapType.Key()

	// Case: key type implements the fmt.Stringer interface
	// Function returns after this block when the condition is true.
	if mapKeyType.Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()) {

		// Populate keys with string values of map keys
		keyValues := reflect.ValueOf(mapToGetKeysFrom).MapKeys()
		for _, keyValue := range keyValues {
			keyValueInterface := keyValue.Interface()
			keys = append(
				keys,
				keyValueInterface.(fmt.Stringer).String(),
			)
		}
		return
	}

	// Case: key type does not implement fmt.Stringer and is not string
	// Function panics when the condition is true.
	if mapKeyType.Kind() != reflect.String {
		panic("key neither implements fmt.Stringer nor is string kind")
	}

	// Case: key is of type string
	keyValues := reflect.ValueOf(mapToGetKeysFrom).MapKeys()
	for _, keyValue := range keyValues {

		// Populate keys with values of map keys
		keyValueInterface := keyValue.Interface()
		keys = append(
			keys,
			keyValueInterface.(string),
		)
	}
	return
}
