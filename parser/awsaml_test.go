package parser

import (
	"strings"
	"testing"
)

func TestParseKeyValueReturnsErrorWhenInputHasInsufficientLength(t *testing.T) {
	_, _, err := ParseKeyValue("")

	if err == nil {
		t.Error("ParseKeyValue did not return error")

	} else if !strings.Contains(err.Error(), "string too short") {
		t.Error("ParseKeyValue returned incorrect message " + err.Error())
	}
}

func TestParseKeyValueReturnsErrorWhenInputDoesNotStartWithSet(t *testing.T) {
	_, _, err := ParseKeyValue("put  ")

	if err == nil {
		t.Error("ParseKeyValue did not return error")

	} else if !strings.Contains(err.Error(), "does not start with 'set '") {
		t.Error("ParseKeyValue returned incorrect message " + err.Error())
	}
}

func TestParseKeyValueDoesNotReturnErrorWhenInputWellFormatted(t *testing.T) {
	_, _, err := ParseKeyValue("set key=value")

	if err != nil {
		t.Error("ParseKeyValue returned unexpected error " + err.Error())
	}

}
