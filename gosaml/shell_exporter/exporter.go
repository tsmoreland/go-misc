package shell_exporter

import (
	"fmt"
	"gosaml/parser"
	"strings"
)

func GetExporter(shell_name string) (func(parser.KeyValuePair) string, error) {
	switch strings.ToUpper(shell_name) {
	case "POWERSHELL":
		return FormatForPowerShell, nil
	default:
		return nil, fmt.Errorf("unsupported shell type")
	}
}
