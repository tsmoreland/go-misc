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

		fmt.Print(line)

		pair, err := parser.ParseKeyValue(line)
		if err != nil {
			fmt.Print(err.Error())
			break
		}

		fmt.Printf("'%s' = '%s'\n", pair.Key(), pair.Value())
	}
}
