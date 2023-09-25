package slice

import (
	"github.com/2hangpeng/lib/compare"
)

// RemoveDuplicate removes duplicate elements from a slice.
func RemoveDuplicate(slice []interface{}) []interface{} {
	if len(slice) == 0 {
		return []interface{}{}
	}
	out := make([]interface{}, 0, len(slice))
	temp := make(map[interface{}]struct{}) // Use struct{} as the value to save memory
	for _, v := range slice {
		if _, ok := temp[v]; !ok {
			temp[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// Contain checks if the slice contains the target element.
// Make sure the element types in the slice and the target are the same.
func Contain(slice []interface{}, target interface{}) bool {
	for _, v := range slice {
		if compare.Compare(v, target) {
			return true
		}
	}
	return false
}

// Chunk divides a slice into multiple sub-slices of the specified size.
func Chunk(slice []interface{}, size int) [][]interface{} {
	var out [][]interface{}
	if len(slice) == 0 || size <= 0 {
		return out
	}
	length := len(slice)
	if size == 1 || size >= length {
		for _, v := range slice {
			out = append(out, []interface{}{v})
		}
		return out
	}
	// Divide the slice evenly
	divideNum := (length + size - 1) / size // Calculate the number of divisions, rounding up
	for i := 0; i < divideNum; i++ {
		start := i * size
		end := (i + 1) * size
		if end > length {
			end = length
		}
		out = append(out, slice[start:end])
	}
	return out
}
