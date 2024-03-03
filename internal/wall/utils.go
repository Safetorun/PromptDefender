package wall

import (
	"time"
)

type RetryableFunc func() (*float64, error)

func Retry(attempts int, sleep time.Duration, fn RetryableFunc) (*float64, error) {
	returnVal, err := fn()

	if err != nil {
		if s, ok := err.(stop); ok {
			return nil, s.error
		}

		if attempts--; attempts > 0 {
			// Add some sleep here if you want to prevent spamming retries
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, fn)
		}

		return nil, err
	}
	return returnVal, nil
}

type stop struct {
	error
}
