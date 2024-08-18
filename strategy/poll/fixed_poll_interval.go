package poll

import "time"

// fixedPollInterval is a poll interval strategy that consistently returns the same fixed interval,
// regardless of the iteration count. This strategy ensures that the polling frequency remains constant
// throughout all iterations, with no backoff or variation.
//
// Example: For iteration: 0 to N and interval: 100 * time.Millisecond
//
//	Iteration 0: 100ms
//	Iteration 1: 100ms
//	Iteration 2: 100ms
//	...
//	Iteration N: 100ms
//
// Result: The interval remains fixed at 100ms for every iteration.
func fixedPollInterval(iteration int, interval time.Duration) time.Duration {
	return interval
}
