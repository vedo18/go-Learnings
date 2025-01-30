package main

import "fmt"

func main() {
	fmt.Println("Hello golang")
	defer fmt.Println("Defer example1")
	defer fmt.Println("Defer example 2")

	fmt.Println("Program end here")
}
