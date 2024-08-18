package utils

import "time"

func MultiplyDuration(iteration int, interval time.Duration) time.Duration {
	return interval * (1 << iteration)
}
