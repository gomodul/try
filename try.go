package try

import (
	"fmt"
)

// MaxRetries is the maximum number of retries
var MaxRetries = 3

// Do keep trying the function until max retry limit or no return error
func Do(fn func(attempt int) error, maxRetries ...int) error {
	var (
		err     error
		attempt = 1
		max     = MaxRetries
	)

	if len(maxRetries) == 1 {
		max = maxRetries[0]
	}

	for {
		(func() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("panic: %v", r)
				}
			}()
			err = fn(attempt)
		})()
		if err == nil {
			break
		}

		attempt++
		if attempt > max {
			return ErrMaxRetriesReached
		}
	}
	return err
}

// IsMaxRetries is a function to check if the error has reached the maximum to try
func IsMaxRetries(err error) bool {
	return err == ErrMaxRetriesReached
}
