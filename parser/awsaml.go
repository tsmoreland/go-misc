package parser

import (
	"fmt"
	"strings"
)

func ParseKeyValue(source string) (key, value string, err error) {

	if len(source) < 5 {
		return "", "", fmt.Errorf("invalid input: string too short")
	}

	if prefix := strings.ToUpper(string(source[0:4])); prefix != "SET " {
		return "", "", fmt.Errorf("invalid format: does not start with 'set '")
	}

	remaining := strings.TrimSpace(string(source[3:]))
	fmt.Print(remaining)

	return "", "", nil
}
