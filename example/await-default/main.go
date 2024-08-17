package main

import (
	"github.com/bayraktugrul/go-await"
	"time"
)

func main() {
	condition := false

	go func() {
		// simulate long running task. (Database insert, network call etc)
		time.Sleep(5 * time.Second)
		condition = true
	}()

	// parameters are not set. Use default values.
	err := await.New().Await(func() bool {
		return condition
	})

	if err != nil {
		// do something
	}
}
