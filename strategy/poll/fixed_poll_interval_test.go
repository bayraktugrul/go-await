package poll_test

import (
	"github.com/bayraktugrul/go-await/strategy/poll"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_FixedInterval(t *testing.T) {
	t.Run("it should return 100ms ms when iteration is 3 and interval is 100ms", func(t *testing.T) {
		//given
		iteration := 3
		interval := 100 * time.Millisecond

		//when
		result := poll.FixedPollInterval(iteration, interval)

		//then
		assert.Equal(t, 100*time.Millisecond, result)
	})
}