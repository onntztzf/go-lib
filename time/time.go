/**
 * @brief
 * @file time
 * @author zhangpeng
 * @version 1.0
 * @date
 */

package time

import (
	"time"
)

//GetInterval Gets the interval of two times
func GetInterval(t1, t2 time.Time) time.Duration {
	duration := t1.Sub(t2)
	if duration < 0 {
		duration = -duration
	}
	return duration
}
