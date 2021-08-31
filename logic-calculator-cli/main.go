package main

import (
	"fmt"
	"logic-calculator-cli/calculator"
	"os"
)

func main() {
	argsWithoutProgram := os.Args[1:]

	var equation string

	if len(argsWithoutProgram) >= 1 {
		equation = argsWithoutProgram[0]
	} else {
		equation = "3 | 2 ^ 1 & 3"
	}

	expected := 3 | 2 ^ 1&3
	fmt.Println(expected)

	// order here is  3 or ((2 xor 1) and 1) =
	op, err := calculator.New(equation)
	if err != nil {
		fmt.Println(err)
		return
	}

	actual, err := op.Solve()
	if err != nil {
		fmt.Println(err)
		return
	}

	if actual != int64(expected) {
		fmt.Printf("%v does not equal %v", actual, expected)
	}

}
