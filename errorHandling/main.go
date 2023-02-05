package main

import (
	"errors"
	"fmt"
	"log"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, NewTimestampedError("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
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
