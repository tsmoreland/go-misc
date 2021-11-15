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
		fmt.Printf("usage %v \"equation here\"\n", os.Args[0])
		return
	}

	// order here is  3 or ((2 xor 1) and 1) =
	op, err := calculator.NewOperation(equation)
	if err != nil {
		fmt.Println(err)
		return
	}

	actual, err := op.Solve()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v", actual)
	}
}
