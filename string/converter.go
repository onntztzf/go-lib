package string

import (
	"strconv"
	"strings"
	"unsafe"
)

func FirstLower(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func FirstUpper(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func Camel2Snake(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// converts case to case through ASCII code
		// 65-90 (A-Z), 97-122 (a-z)
		// if the letter A-Z is capitalized, put an _ in front of it
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//convert all uppercase letters to lowercase letters
	return strings.ToLower(string(data[:]))
}

func Snake2Camel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func ToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// ToBool convert string to a boolean
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// ToChar convert string to char slice
func ToChar(s string) []string {
	c := make([]string, 0)
	if len(s) == 0 {
		c = append(c, "")
	}
	for _, v := range s {
		c = append(c, string(v))
	}
	return c
}

func FromBytes(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
