package string

import "regexp"

func Match(s string, pattern string) (bool, error) {
	if len(s) == 0 && len(pattern) == 0 {
		return true, nil
	}
	match, err := regexp.MatchString(pattern, s)
	return match, err
}
