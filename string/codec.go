package string

import (
	"encoding/base64"
)

// EncodeByBase64 encodes a string using base64 encoding.
func EncodeByBase64(target string) string {
	strBytes := []byte(target)
	return base64.StdEncoding.EncodeToString(strBytes)
}

// DecodeByBase64 decodes a base64-encoded string and returns the original string.
func DecodeByBase64(target string) (string, error) {
	result, err := base64.StdEncoding.DecodeString(target)
	if err != nil {
		return string(result), err
	}
	return string(result), nil
}
