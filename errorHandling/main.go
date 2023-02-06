package main

import (
	"context"
)

func main() {
	DemoTimestampedError()

	files := []string{
		"main.go",
		"timestamperror.go",
		"404.go",
	}

	BadWaitGroup(files)
	ErrGroup(files)
	ErrGroupWithContext(context.Background(), files)
}
