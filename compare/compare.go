package compare

import "reflect"

// Compare two interfaces
func Compare(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
