package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("I will be executed when the function ends, So I am executed at the end")
	fmt.Println("Sleeping for 1 second ...")
	time.Sleep(time.Second)
	fmt.Println("Time's up!")
}
