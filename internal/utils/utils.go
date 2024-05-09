package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
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
			time.Sleep(sleep)
			log.Default().Println("Retrying after error: ", err, " Attempts left: ", attempts)
			return Retry(attempts, 2*sleep, fn)
		}

		return nil, err
	}
	return returnVal, nil
}

func HashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

type stop struct {
	error
}
