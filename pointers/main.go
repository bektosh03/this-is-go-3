package main

import "fmt"

func main() {
	// 1:
	a := 5
	a = changeValueWithReturn(a)
	fmt.Println("a:", a)
	// 2:
	b := 3
	changeValueWithPointer(&b)
	fmt.Println("b:", b)
}

func changeValueWithReturn(num int) int {
	num = 10
	return num
}

func changeValueWithPointer(num *int) {
	*num = 10
}
