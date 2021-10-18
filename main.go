package main

import (
	"bufio"
	"fmt"
	"gosaml/parser"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			break
		}

		pair, err := parser.ParseKeyValue(line)
		if err != nil {
			fmt.Print(err.Error())
			break
		}

		if pair.Equals(parser.EmptyKeyValuePair) {
			continue
		}

		key, value := pair.Deconstruct()

		fmt.Printf("$env:%s='%s'", key, value)

	}
}
