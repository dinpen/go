package tasks

import (
	"errors"
	"log"
	"testing"
)

func TestRetry(t *testing.T) {
	RetryOnFalse(5, func() bool {
		log.Println("return true")
		return true
	})
	RetryOnFalse(5, func() bool {
		log.Println("return false")
		return false
	})

	RetryOnError(5, func() error {
		log.Println("return nil")
		return nil
	})
	RetryOnError(5, func() error {
		log.Println("return error")
		return errors.New("err")
	})
}
