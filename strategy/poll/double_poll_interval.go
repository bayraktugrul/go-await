package poll

import (
	"github.com/bayraktugrul/go-await/utils"
	"time"
)

// DoublePollInterval TODO:write explanation
func DoublePollInterval(iteration int, interval time.Duration) time.Duration {
	return utils.MultiplyDuration(iteration, interval)
}
