package string

import (
	"regexp"
	"sync"
)

var (
	regexCache     = make(map[string]*regexp.Regexp)
	regexCacheLock sync.RWMutex
)

// Match checks if a string matches a given pattern using a cached compiled regular expression.
func Match(s string, pattern string) (bool, error) {
	if len(s) == 0 && len(pattern) == 0 {
		return true, nil
	}

	// Check if the compiled regular expression is already in the cache.
	regexCacheLock.RLock()
	re, found := regexCache[pattern]
	regexCacheLock.RUnlock()

	if !found {
		// If not found, compile the regular expression and store it in the cache.
		var err error
		re, err = regexp.Compile(pattern)
		if err != nil {
			return false, err
		}

		regexCacheLock.Lock()
		regexCache[pattern] = re
		regexCacheLock.Unlock()
	}

	return re.MatchString(s), nil
}
