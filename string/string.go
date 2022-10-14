package string

import (
	"github.com/pkg/errors"
	"regexp"
)

var (
	// ErrInvalidStartPosition is an error that indicates the start position is invalid.
	ErrInvalidStartPosition = errors.New("start position is invalid")
	// ErrInvalidStopPosition is an error that indicates the stop position is invalid.
	ErrInvalidStopPosition = errors.New("stop position is invalid")
)

func Match(pattern string, s string) (bool, error) {
	if len(pattern) == 0 && len(s) == 0 {
		return true, nil
	}
	match, err := regexp.MatchString(pattern, s)
	return match, err
}

// Contains checks if str is in list.
func Contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}
	return false
}

// Filter filters chars from s with given filter function.
func Filter(s string, filter func(r rune) bool) string {
	var n int
	chars := []rune(s)
	for i, x := range chars {
		if n < i {
			chars[n] = x
		}
		if !filter(x) {
			n++
		}
	}
	return string(chars[:n])
}

// Remove removes given strs from strings.
func Remove(strings []string, strs ...string) []string {
	out := append([]string(nil), strings...)
	for _, str := range strs {
		var n int
		for _, v := range out {
			if v != str {
				out[n] = v
				n++
			}
		}
		out = out[:n]
	}
	return out
}

func Substr(str string, start, stop int) (string, error) {
	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		return "", ErrInvalidStartPosition
	}
	if stop < 0 || stop > length {
		return "", ErrInvalidStopPosition
	}
	return string(rs[start:stop]), nil
}
