package shell_exporter

import (
	"fmt"
	"gosaml/parser"
)

func FormatForPowerShell(pair parser.KeyValuePair) string {
	key, value := pair.Deconstruct()
	return fmt.Sprintf("$env:%s='%s'", key, value)
}
