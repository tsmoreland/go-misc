package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || len(line) == 0 {
			break
		}
		fmt.Printf("%s %d", line, len(line))
	}
}
