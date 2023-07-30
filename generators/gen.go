// go:build ignore
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := os.Getenv("GOFILE")
	fmt.Printf("Running generate on %s\n", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var count int
	for scanner.Scan() {
		count++
	}
	fmt.Printf("Number of lines in file: %d", count)
}
