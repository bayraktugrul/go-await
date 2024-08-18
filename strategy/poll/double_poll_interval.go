package poll

import (
	"github.com/bayraktugrul/go-await/utils"
	"time"
)

// DoublePollInterval calculates the poll interval for a given iteration
// using an exponential backoff strategy. The interval is doubled with each iteration,
// starting from the initial interval value provided.
//
// Iterations start from 0 (zero). At each step, the interval is multiplied by 2.
//
// Example1: For iteration: 3 and interval: 100 * time.Millisecond
//
//	Step 1: 100ms * 2 = 200ms
//	Step 2: 200ms * 2 = 400ms
//	Step 3: 400ms * 2 = 800ms
//
// Result: 800ms
func DoublePollInterval(iteration int, interval time.Duration) time.Duration {
	return utils.MultiplyDuration(iteration, interval)
}
