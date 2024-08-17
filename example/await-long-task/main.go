package main

import (
	"github.com/bayraktugrul/go-await"
	"time"
)

func main() {
	condition := false

	go func() {
		// simulate long running task. (Database insert, network call etc)
		time.Sleep(60 * time.Second)
		condition = true
	}()

	err := await.New().AtMost(2 * time.Second).PollInterval(200 * time.Millisecond).Await(func() bool {
		return condition
	})

	// returns error because condition is not met in 2 seconds.
	if err != nil {
		// do something
	}
}