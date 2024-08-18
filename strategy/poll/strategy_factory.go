package poll

import (
	"time"
)

func Strategies() map[Strategy]func(int, time.Duration) time.Duration {
	return map[Strategy]func(int, time.Duration) time.Duration{
		Fixed:  fixedPollInterval,
		Double: doublePollInterval,
	}
}
