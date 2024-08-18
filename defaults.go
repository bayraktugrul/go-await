package await

import (
	"github.com/bayraktugrul/go-await/strategy/poll"
	"time"
)

const (
	DefaultTimeout      = 1 * time.Second
	DefaultInterval     = 100 * time.Millisecond
	DefaultPollStrategy = poll.Fixed
)
