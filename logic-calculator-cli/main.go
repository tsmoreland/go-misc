package main

import (
	"fmt"
	"logic-calculator-cli/calculator"
)

func main() {
	fmt.Println("Pending...")

	expected := 3 | 2 ^ 1&3
	fmt.Println(expected)

	// order here is  3 or ((2 xor 1) and 1) =
	op, err := calculator.NewOperation("3 | 2 ^ 1 & 3")
	if err != nil {
		fmt.Println(err)
		return
	}

	actual := op.Solve()

	if actual != int64(expected) {
		fmt.Printf("%v does not equal %v", actual, expected)
	}

}
