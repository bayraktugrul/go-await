package utils_test

import (
	"github.com/bayraktugrul/go-await/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Util_MultiplyDuration(t *testing.T) {
	//given
	cases := []struct {
		definition string
		iteration  int
		interval   time.Duration
		result     time.Duration
	}{
		{
			definition: "it should return 100 ms when interval is 100ms and iteration is 0",
			iteration:  0,
			interval:   100 * time.Millisecond,
			result:     100 * time.Millisecond,
		},
		{
			definition: "it should return 200 ms when interval is 100ms and iteration is 1",
			iteration:  1,
			interval:   100 * time.Millisecond,
			result:     200 * time.Millisecond,
		},
		{
			definition: "it should return 400 ms when interval is 100ms and iteration is 2",
			iteration:  2,
			interval:   100 * time.Millisecond,
			result:     400 * time.Millisecond,
		},
		{
			definition: "it should return 800 ms when interval is 100ms and iteration is 3",
			iteration:  3,
			interval:   100 * time.Millisecond,
			result:     800 * time.Millisecond,
		},

		{
			definition: "it should return 8s when interval is 1s and iteration is 2",
			iteration:  3,
			interval:   1 * time.Second,
			result:     8 * time.Second,
		},
	}

	for _, c := range cases {
		t.Run(c.definition, func(t *testing.T) {
			result := utils.MultiplyDuration(c.iteration, c.interval)
			assert.Equal(t, c.result, result)
		})
	}
}
