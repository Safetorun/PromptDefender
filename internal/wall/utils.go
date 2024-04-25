package wall

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/safetorun/PromptDefender/cache"
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

func checkCache(cache *cache.Cache, prompt PromptToCheck) (bool, *CheckResult, error) {
	if cache == nil {
		return false, nil, nil
	}

	b, err := json.Marshal(prompt)

	if err != nil {
		return false, nil, err
	}

	cachedResult, err := (*cache).Get(HashString(string(b)))

	if err != nil {
		return false, nil, err
	}

	if cachedResult != nil {
		var cachedResultReturn *CheckResult
		err := json.Unmarshal([]byte(*cachedResult), &cachedResultReturn)

		if err != nil {
			return false, nil, err
		}

		return true, cachedResultReturn, nil
	}

	return false, nil, nil
}

func storeCache(cache *cache.Cache, prompt PromptToCheck, result *CheckResult) error {
	if cache == nil {
		return nil
	}

	b, err := json.Marshal(prompt)

	if err != nil {
		return err
	}

	bResult, err := json.Marshal(result)

	if err != nil {
		return err
	}

	return (*cache).Set(HashString(string(b)), string(bResult))
}

type stop struct {
	error
}
