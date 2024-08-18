package await_test

import (
	"errors"
	"github.com/bayraktugrul/go-await"
	"github.com/bayraktugrul/go-await/strategy/poll"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_App_Error(t *testing.T) {
	//given
	cases := []struct {
		definition     string
		actualInterval time.Duration
		actualTimeout  time.Duration
		waitFunc       func() bool
		result         error
	}{
		{
			definition:     "it should return error when task is not finished in 500 ms timeout",
			actualInterval: 100 * time.Millisecond,
			actualTimeout:  500 * time.Millisecond,
			waitFunc: func() bool {
				time.Sleep(2 * time.Second)
				return false
			},
			result: errors.New("condition does not match in 500ms timeout time"),
		},
		{
			definition:     "it should return error when task is not finished in 500 ms timeout",
			actualInterval: 100 * time.Millisecond,
			actualTimeout:  500 * time.Millisecond,
			waitFunc:       nil,
			result:         errors.New("function can not be nil"),
		},
		{
			definition:     "it should return error when poll interval is zero",
			actualInterval: 0,
			actualTimeout:  500 * time.Millisecond,
			waitFunc: func() bool {
				time.Sleep(2 * time.Second)
				return false
			},
			result: errors.New("poll interval must be greater than zero. internal: 0s"),
		},
		{
			definition:     "it should return error when timeout is zero",
			actualInterval: 100 * time.Millisecond,
			actualTimeout:  0,
			waitFunc: func() bool {
				time.Sleep(2 * time.Second)
				return false
			},
			result: errors.New("timeout must be greater than zero. timeout: 0s"),
		},
		{
			definition:     "it should return error when poll interval is greater than timeout",
			actualInterval: 100 * time.Millisecond,
			actualTimeout:  50 * time.Millisecond,
			waitFunc: func() bool {
				time.Sleep(2 * time.Second)
				return false
			},
			result: errors.New("poll interval can not be greater than timeout. timeout: 50ms interval: 100ms"),
		},
	}

	for _, c := range cases {
		t.Run(c.definition, func(t *testing.T) {
			app := await.New()
			result := app.AtMost(c.actualTimeout).PollInterval(c.actualInterval).Await(c.waitFunc)

			assert.Equal(t, c.result, result)
		})
	}
}

func Test_App_HappyPath(t *testing.T) {
	//given
	t.Run("it should wait 1 second task with 2 seconds timeout and 100ms poll intervals", func(t *testing.T) {

		condition := false
		go func() {
			time.Sleep(1 * time.Second)
			condition = true
		}()

		err := await.New().AtMost(2 * time.Second).PollInterval(100 * time.Millisecond).Await(func() bool {
			return condition
		})

		assert.NoError(t, err)
	})

	t.Run("it should wait 1 seconds task with 2s timeout and 100ms poll intervals with double poll strategy", func(t *testing.T) {

		condition := false
		go func() {
			time.Sleep(1 * time.Second)
			condition = true
		}()

		err := await.New().PollStrategy(poll.Double).
			PollInterval(100 * time.Millisecond).
			AtMost(2 * time.Second).
			Await(func() bool {
				return condition
			})

		assert.NoError(t, err)
	})
}
