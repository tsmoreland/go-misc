package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// TimestampedError a simple error containing both a message and a timestamp
type TimestampedError struct {
	timestamp time.Time
	message   string
}

// NewTimestampedError creates a new error instance using the provided message
//
//	and current time as the timestamp
func NewTimestampedError(message string) error {
	return &TimestampedError{
		timestamp: time.Now(),
		message:   message,
	}
}

func (e TimestampedError) Error() string {
	time := e.timestamp.Format("2006-01-02T15:04:05.999999999Z07:00")
	return fmt.Sprintf("[%s]: %s", time, e.message)
}

// Equal compares the message content of the current error with other
//
//	can be used for equality checks with go-cmp
func (e TimestampedError) Equal(other TimestampedError) bool {
	return e.message == other.message
}

// DemoTimestampedError demonstrates use of timestamped error handling
func DemoTimestampedError() {
	divide := func(a int, b int) (int, error) {
		if b == 0 {
			return 0, NewTimestampedError("cannot divide by zero")
		}
		return a / b, nil
	}
	result, err := divide(1, 0)
	var timestampedError *TimestampedError
	if errors.As(err, &timestampedError) {
		log.Fatalf("error with time: %v", timestampedError)
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(result)
	}
}
