package main

import "fmt"

func main() {
	fmt.Println("Hello golang")

	// Schedule "Defer example1" to be printed just before the main function returns.
	// The defer statement ensures this line will execute after all other code in the main function.
	defer fmt.Println("Defer example1")

	// Schedule "Defer example 2" to be printed just before the main function returns.
	// This defer statement is added to the stack of deferred calls.
	defer fmt.Println("Defer example 2")

	fmt.Println("Program end here")

	// At this point, the main function is about to return.
	// The deferred functions are executed in Last-In-First-Out (LIFO) order:
	// 1. "Defer example 2" (last deferred) is executed first.
	// 2. "Defer example1" (first deferred) is executed next.
}
