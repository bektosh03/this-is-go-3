package main

import (
	"fmt"
	"strconv"
)

type operationFunction func(int, int) int

var opMap = map[string]operationFunction{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func Calculate(exp []string, index int) int {
	if len(exp) < 3 {
		fmt.Printf("the expression at index %d is not valid\n", index)
		return 0
	}
	a, err := strconv.Atoi(exp[0])
	if err != nil {
		fmt.Printf("the first parameter of the expression at index %d is not valid\n", index)
		return 0
	}
	op := exp[1]
	opFunc, ok := opMap[op]
	if !ok {
		fmt.Printf("operator in the expression at index %d is not supported\n", index)
		return 0
	}
	b, err := strconv.Atoi(exp[2])
	if err != nil {
		fmt.Printf("the second parameter of the expression at index %d is not valid\n", index)
		return 0
	}
	result := opFunc(a, b)
	fmt.Printf("The result of expression at index %d is %d\n", index, result)
	return result
}

func main() {
	expressions := [][]string{
		[]string{"1", "+", "2"},
		[]string{"3", "-", "1"},
		[]string{"3", "*", "2"},
		[]string{"6", "/", "2"},
		[]string{"two", "+", "3"},
		[]string{"2", "+", "three"},
		[]string{"8"},
	}

	for index, exp := range expressions {
		Calculate(exp, index + 1)
	}
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
