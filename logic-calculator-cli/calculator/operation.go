package calculator

import (
	"errors"
	"strconv"
	"strings"
)

type Operation struct {
	left      *Operation
	operation string
	right     *Operation
	value     int64
}

func NewOperation(source string) (*Operation, error) {
	op := &Operation{value: 0}

	return op.parse(source)
}

func (op *Operation) parse(source string) (*Operation, error) {
	// todo: investigate to see if we can use slice and keep string functionality (pass as *string to minimize copy)
	// order of operations = ~ ^ & |  -> NOT XOR AND OR, ~ can be accomplished by 1^M where M is the value but that doesn't involve left/right so may have to be done in place

	source = strings.Replace(source, " ", "", -1)

	// todo change to for loop over the operator types, after the loop try to see if it's a number

	operators := []string{"^", "&", "|"}
	for _, operator := range operators {
		found, err := op.process(source, operator)
		if found {
			return op, nil
		} else if err != nil {
			return nil, err
		}
	}

	value, err := strconv.ParseInt(source, 10, 64)
	if err != nil {
		return nil, err
	}

	op.value = value
	op.left = nil
	op.right = nil
	op.operation = "value"

	return op, nil
}

func numberOfSides(operator string) int16 {
	switch operator {
	case "~":
		return 1
	default:
		return 2
	}
}

func (op *Operation) process(source string, operator string) (bool, error) {
	index := strings.LastIndex(source, operator)
	if index == -1 {
		return false, nil
	}

	switch numberOfSides(operator) {
	case 2:
		left, right, err := split(index, source)
		if err != nil {
			return false, err
		}
		op.left = left
		op.right = right
		op.operation = operator
	case 1:
		op.left = nil
		op.right = nil
		// ??? we need something like split here, first pass will only support something like
		// ~x | y  I'd expect ~x to be treated as a result like x on its own (i.e. calculate the result in place)
		// that result lets say X would then be the left of X | y
		// the question is how to we get just ~x and not get, we need something like split (maybe even split itself)
		// but where we expct left to come back as nil
		return false, errors.New("not implemented")

	default:
		return true, errors.New("unrecognized equation")
	}

	return true, nil
}

func split(index int, source string) (*Operation, *Operation, error) {
	if index == -1 {
		return nil, nil, errors.New("invalid index")
	}
	if index+1 >= len(source) {
		return nil, nil, errors.New("index out of range, malformed operation")
	}

	var left *Operation
	var err error
	left, err = NewOperation(source[0:index])
	if err != nil {
		return nil, nil, err
	}

	var right *Operation
	right, err = NewOperation(source[index+1:])
	if err != nil {
		return nil, nil, err
	}

	return left, right, nil
}

func (op *Operation) Solve() (int64, error) {

	if op.operation == "value" {
		return op.value, nil
	}

	if op.left == nil {
		return 0, errors.New("left operation undefined")
	}
	if op.right == nil {
		return 0, errors.New("rigth operation undefined")
	}

	value, error := op.left.Solve()
	if error != nil {
		return 0, error
	}
	leftValue := value
	value, error = op.right.Solve()
	if error != nil {
		return 0, error
	}

	switch op.operation {
	case "~":
		// can be achieved with 1 ^ ??? doesn't require left and right though
		return 0, errors.New("not implemented yet")
	case "^":
		return leftValue ^ value, nil
	case "|":
		return leftValue | value, nil
	case "&":
		return leftValue & value, nil
	default:
		return 0, errors.New("undefined operation" + op.operation)
	}
}
