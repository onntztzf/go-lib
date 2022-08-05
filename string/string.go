/**
 * @brief
 * @file string
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package string

import (
	"regexp"
	"strings"
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

func Match(pattern string, s string) (bool, error) {
	if len(pattern) == 0 && len(s) == 0 {
		return true, nil
	}
	match, err := regexp.MatchString(pattern, s)
	return match, err
}
