package await

import (
	"fmt"
	"time"
)

const (
	DefaultTimeout  = 1 * time.Second
	DefaultInterval = 100 * time.Millisecond
)

type app struct {
	timeout  time.Duration
	interval time.Duration
}

type App interface {
	AtMost(timeout time.Duration) App
	PollInterval(interval time.Duration) App
	Await(waitFunc func() bool) error
}

func New() App {
	return &app{
		timeout:  DefaultTimeout,
		interval: DefaultInterval,
	}
}

func (a *app) AtMost(timeout time.Duration) App {
	a.timeout = timeout
	return a
}

func (a *app) PollInterval(interval time.Duration) App {
	a.interval = interval
	return a
}

func (a *app) Await(waitFunc func() bool) error {
	if waitFunc == nil {
		return fmt.Errorf("function can not be nil")
	}
	err := a.validateDuration()
	if err != nil {
		return err
	}

	start := time.Now()
	remaining := a.timeout
	result := make(chan bool)

	go until(waitFunc, result)

	for {
		select {
		case finish := <-result:
			if finish {
				return nil
			}
			remaining = a.timeout - time.Now().Sub(start)
			if remaining <= 0 {
				return fmt.Errorf("condition does not match in %s timeout time", a.timeout)
			}
			go until(waitFunc, result)
		case <-time.After(remaining):
			return fmt.Errorf("condition does not match in %s timeout time", a.timeout)
		}

		time.Sleep(a.interval)
	}
}

func until(until func() bool, result chan bool) {
	result <- until()
}

func (a *app) validateDuration() error {
	if a.interval <= 0 {
		return fmt.Errorf("poll interval must be greater than zero. internal: %s", a.interval)
	}
	if a.timeout <= 0 {
		return fmt.Errorf("timeout must be greater than zero. timeout: %s", a.timeout)
	}

	if a.interval > a.timeout {
		return fmt.Errorf("poll interval can not be greater than timeout. timeout: %s interval: %s", a.timeout, a.interval)
	}
	return nil
}
