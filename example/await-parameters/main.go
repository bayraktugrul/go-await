package main

import (
	"github.com/bayraktugrul/go-await"
	"time"
)

func main() {
	condition := false

	go func() {
		// simulate long running task. (Database insert, network call etc)
		time.Sleep(200 * time.Millisecond)
		condition = true
	}()

	err := await.New().PollInterval(200 * time.Millisecond).AtMost(10 * time.Second).Await(func() bool {
		return condition
	})

	if err != nil {
		// do something
	}
}
