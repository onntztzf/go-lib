package slice

import (
	"github.com/gh-zhangpeng/box-lib/compare"
)

//RemoveDuplicate Delete duplicate elements
func RemoveDuplicate(slice []interface{}) []interface{} {
	if len(slice) == 0 {
		return []interface{}{}
	}
	out := make([]interface{}, 0, len(slice))
	temp := make(map[interface{}]bool)
	for i := range slice {
		if !temp[slice[i]] {
			temp[slice[i]] = true
			out = append(out, slice[i])
		}
	}
	return out
}

//Contain check if the value is in the iterable type or not
//You need to ensure that the target type is the same as the element type in slice
func Contain(slice []interface{}, target interface{}) bool {
	for _, v := range slice {
		if compare.Compare(v, target) {
			return true
		}
	}
	return false
}

//Chunk creates a slice of elements splits into groups the length of `size`
func Chunk(slice []interface{}, size int) [][]interface{} {
	var out [][]interface{}
	if len(slice) == 0 || size <= 0 {
		return out
	}
	length := len(slice)
	if size == 1 || size >= length {
		for _, v := range slice {
			var tmp []interface{}
			tmp = append(tmp, v)
			out = append(out, tmp)
		}
		return out
	}
	// divide slice equally
	divideNum := length/size + 1
	for i := 0; i < divideNum; i++ {
		if i == divideNum-1 {
			if len(slice[i*size:]) > 0 {
				out = append(out, slice[i*size:])
			}
		} else {
			out = append(out, slice[i*size:(i+1)*size])
		}
	}
	return out
}
