package parser

import (
	"fmt"
	"regexp"
	"strings"
)

func ParseKeyValue(source string) (pair KeyValuePair, err error) {

	if len(source) < 5 {
		return EmptyKeyValuePair, fmt.Errorf("invalid input: string too short")
	}

	if prefix := strings.ToUpper(string(source[0:4])); prefix != "SET " {
		return EmptyKeyValuePair, fmt.Errorf("invalid format: does not start with 'set '")
	}

	remaining := strings.TrimSpace(string(source[3:]))

	equals := regexp.MustCompile("=")
	parts := equals.Split(remaining, 2)

	if len(parts) != 2 {
		return EmptyKeyValuePair, fmt.Errorf("invalid format: unable to split key and value, missing '='")
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	if len(key) == 0 {
		return EmptyKeyValuePair, fmt.Errorf("invalid format: key cannot be empty")
	}

	return KeyValuePair{key, value}, nil
}
