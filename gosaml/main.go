package main

import (
	"bufio"
	"fmt"
	"gosaml/parser"
	"gosaml/shell_exporter"
	"os"
	"strings"
)

func main() {

	formatter, err := shell_exporter.GetExporter("powershell")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	reader := bufio.NewReader(os.Stdin)

	commands := []string{}

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

		command := formatter(pair)
		commands = append(commands, command)
	}

	for _, command := range commands {
		fmt.Printf(command + "\n")
	}
}
