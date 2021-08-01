package main

import "fmt"

func main() {
	// TODO 8: anonymous functions
	x := func() int {
		fmt.Println("this is an anonymous function")
		return 5
	}()
	fmt.Println(x)
}