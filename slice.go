/**
 * @brief
 * @file slice
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package box_lib

import (
	"fmt"
	"reflect"
)

//sliceValue return the reflect value of a slice
func sliceValue(slice interface{}) reflect.Value {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Invalid slice type, value of type %T", slice))
	}
	return v
}

//Unique remove duplicate elements in slice
func Unique(slice interface{}) interface{} {
	sv := sliceValue(slice)
	if sv.Len() == 0 {
		return slice
	}
	var temp []interface{}
	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i).Interface()
		skip := true
		for j := range temp {
			if v == temp[j] {
				skip = false
				break
			}
		}
		if skip {
			temp = append(temp, v)
		}
	}
	res := reflect.MakeSlice(sv.Type(), len(temp), len(temp))
	for i := 0; i < len(temp); i++ {
		res.Index(i).Set(reflect.ValueOf(temp[i]))
	}
	return res.Interface()
}

// StringSlice convert param to slice of string
func StringSlice(slice interface{}) []string {
	v := sliceValue(slice)
	out := make([]string, v.Len())
	for i := 0; i < v.Len(); i++ {
		v, ok := v.Index(i).Interface().(string)
		if !ok {
			panic("invalid element type")
		}
		out[i] = v
	}
	return out
}

// Int64Slice convert param to slice of int64
func Int64Slice(slice interface{}) []int64 {
	sv := sliceValue(slice)
	out := make([]int64, sv.Len())
	for i := 0; i < sv.Len(); i++ {
		v, ok := sv.Index(i).Interface().(int64)
		if !ok {
			panic("invalid element type")
		}
		out[i] = v
	}
	return out
}

func UniqueUIntSlice(arr []uint) []uint {
	occurred := map[uint]struct{}{}
	var result []uint
	for _, value := range arr {
		// check if already the mapped
		// variable is set to true or not
		if _, ok := occurred[value]; !ok {
			occurred[value] = struct{}{}
			// Append to result slice.
			result = append(result, value)
		}
	}
	return result
}

func UniqueInt64Slice(arr []int64) []int64 {
	occurred := map[int64]struct{}{}
	var result []int64
	for _, value := range arr {
		// check if already the mapped
		// variable is set to true or not
		if _, ok := occurred[value]; !ok {
			occurred[value] = struct{}{}
			// Append to result slice.
			result = append(result, value)
		}
	}
	return result
}

func UniqueStringSlice(arr []string) []string {
	occurred := map[string]struct{}{}
	var result []string
	for _, value := range arr {
		// check if already the mapped
		// variable is set to true or not
		if _, ok := occurred[value]; !ok {
			occurred[value] = struct{}{}
			// Append to result slice.
			result = append(result, value)
		}
	}
	return result
}

//Contain check if the value is in the iterable type or not
func Contain(slice interface{}, target interface{}) bool {
	v := reflect.ValueOf(slice)
	switch kind := reflect.TypeOf(slice).Kind(); kind {
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == target {
				return true
			}
		}
	default:
		panic(fmt.Sprintf("kind %s is not support", slice))
	}
	return false
}

//Chunk creates an slice of elements split into groups the length of `size`
func Chunk(slice []interface{}, size int) [][]interface{} {
	var res [][]interface{}
	if len(slice) == 0 || size <= 0 {
		return res
	}
	length := len(slice)
	if size == 1 || size >= length {
		for _, v := range slice {
			var tmp []interface{}
			tmp = append(tmp, v)
			res = append(res, tmp)
		}
		return res
	}
	// divide slice equally
	divideNum := length/size + 1
	for i := 0; i < divideNum; i++ {
		if i == divideNum-1 {
			if len(slice[i*size:]) > 0 {
				res = append(res, slice[i*size:])
			}
		} else {
			res = append(res, slice[i*size:(i+1)*size])
		}
	}
	return res
}

// Difference creates an slice of whose element in slice1, not in slice2
func Difference(slice1, slice2 interface{}) interface{} {
	sv := sliceValue(slice1)
	var indexes []int
	for i := 0; i < sv.Len(); i++ {
		item := sv.Index(i).Interface()
		if !Contain(slice2, item) {
			indexes = append(indexes, i)
		}
	}
	res := reflect.MakeSlice(sv.Type(), len(indexes), len(indexes))
	for i := range indexes {
		res.Index(i).Set(sv.Index(indexes[i]))
	}
	return res.Interface()
}

// Union creates a slice of unique values, in order, from all given slices. using == for equality comparisons
func Union(slices ...interface{}) interface{} {
	if len(slices) == 0 {
		return nil
	}
	// append all slices, then unique it
	var allSlices []interface{}
	l := 0
	for i := range slices {
		sv := sliceValue(slices[i])
		l += sv.Len()
		for j := 0; j < sv.Len(); j++ {
			v := sv.Index(j).Interface()
			allSlices = append(allSlices, v)
		}
	}
	sv := sliceValue(slices[0])
	res := reflect.MakeSlice(sv.Type(), l, l)
	for i := 0; i < l; i++ {
		res.Index(i).Set(reflect.ValueOf(allSlices[i]))
	}
	return Unique(res.Interface())
}
