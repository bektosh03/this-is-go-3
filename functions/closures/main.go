package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	func() {
		fmt.Printf("Hello %s\n", name)
	}()
}
