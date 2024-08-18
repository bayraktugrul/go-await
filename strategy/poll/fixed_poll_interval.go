package poll

import "time"

// FixedPollInterval poll interval strategy always returns given interval
func FixedPollInterval(iteration int, interval time.Duration) time.Duration {
	return interval
}
