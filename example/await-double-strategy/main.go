package main

import (
	"github.com/bayraktugrul/go-await"
	"github.com/bayraktugrul/go-await/strategy/poll"
	"time"
)

func main() {
	condition := false

	go func() {
		// simulate long running task. (Database insert, network call etc)
		time.Sleep(400 * time.Millisecond)
		condition = true
	}()

	// parameters are not set. Use default values.
	err := await.New().
		PollStrategy(poll.Double).
		PollInterval(100 * time.Millisecond).
		AtMost(2 * time.Second).
		Await(func() bool {
			return condition
		})

	if err != nil {
		// do something
	}
}
