package calculator

import (
	"testing"
)

func TestNewReturnsErrorWhenSourceIsEmpty(t *testing.T) {
	_, err := NewOperation("")

	if err == nil {
		t.Error("error was nil")
	}
}

func TestNewReturnsNilErrorWhenSourceIsValid(t *testing.T) {
	_, err := NewOperation("3|2")

	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}
}
