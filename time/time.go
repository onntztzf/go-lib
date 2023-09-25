package time

import (
	"time"
)

// GetInterval returns the interval between two times.
func GetInterval(t1, t2 time.Time) time.Duration {
	duration := t1.Sub(t2)
	if duration < 0 {
		duration = -duration
	}
	return duration
}
