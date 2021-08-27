package calculator

import (
	"errors"
	"strconv"
	"strings"
)

type Operation struct {
	left  *Operation
	value string
	right *Operation
}

func NewOperation(source string) (*Operation, error) {
	op := &Operation{}

	return op.parse(source)
}

func (op *Operation) parse(source string) (*Operation, error) {
	// todo: investigate to see if we can use slice and keep string functionality (pass as *string to minimize copy)
	// order of operations = ! ^ & |  -> NOT XOR AND OR

	source = strings.Replace(source, " ", "", -1)

	// todo change to for loop over the operator types, after the loop try to see if it's a number

	found, err := op.process(source, "^")
	if found {
		return op, nil
	} else if err != nil {
		return nil, err
	}
	found, err = op.process(source, "&")
	if found {
		return op, nil
	} else if err != nil {
		return nil, err
	}

	found, err = op.process(source, "|")
	if found {
		return op, nil
	} else if err != nil {
		return nil, err
	}

	_, err = strconv.ParseInt(source, 10, 64)
	if err != nil {
		return nil, err
	}

	op.left = nil
	op.right = nil
	op.value = source

	return op, nil
}

func (op *Operation) process(source string, operator string) (bool, error) {
	index := strings.LastIndex(source, operator)
	if index != -1 {
		left, right, err := split(index, source)
		if err != nil {
			return false, err
		}
		op.left = left
		op.right = right
		op.value = operator
		return true, nil
	}
	return false, nil
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

func (op *Operation) Solve() int64 {
	return 0
}
