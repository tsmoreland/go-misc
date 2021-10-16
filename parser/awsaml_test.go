package parser

import (
	"strings"
	"testing"
)

func TestParseKeyValueReturnsErrorWhenInputHasInsufficientLength(t *testing.T) {
	_, err := ParseKeyValue("")

	if err == nil {
		t.Error("ParseKeyValue did not return error")

	} else if !strings.Contains(err.Error(), "string too short") {
		t.Error("ParseKeyValue returned incorrect message " + err.Error())
	}
}

func TestParseKeyValueReturnsErrorWhenInputDoesNotStartWithSet(t *testing.T) {
	_, err := ParseKeyValue("put  ")

	if err == nil {
		t.Error("ParseKeyValue did not return error")

	} else if !strings.Contains(err.Error(), "does not start with 'set '") {
		t.Error("ParseKeyValue returned incorrect message " + err.Error())
	}
}

func TestParseKeyValueReturnsErrorWhenInputDoesNotContainEquals(t *testing.T) {
	_, err := ParseKeyValue("set invalid")
	if err == nil {
		t.Error("ParseKeyValue did not return error")

	} else if !strings.Contains(err.Error(), "invalid format: unable to split key and value, missing '='") {
		t.Error("ParseKeyValue returned incorrect message " + err.Error())
	}
}

func TestParseKeyValueDoesNotReturnErrorWhenInputWellFormatted(t *testing.T) {
	_, err := ParseKeyValue("set key=value")

	if err != nil {
		t.Error("ParseKeyValue returned unexpected error " + err.Error())
	}
}

func TestParseKeyValueReturnsExpectedKeyWhenInputWellFormatted(t *testing.T) {
	pair, err := ParseKeyValue("set key=value")

	if err != nil {
		t.Error("ParseKeyValue returned unexpected error " + err.Error())
	}

	if pair.key != "key" {
		t.Error("key does not match expected value " + pair.key)
	}
	if pair.value != "value" {
		t.Error("key does not match expected value " + pair.value)
	}
}
