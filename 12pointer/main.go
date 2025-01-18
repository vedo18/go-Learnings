package main

import "fmt"

func updateValue(num *int) {
	*num = *num + 20
}

func main() {
	// Declare a normal variable
	number := 42

	// Get the pointer to the variable
	ptr := &number

	// Print the value and the address
	fmt.Println("Value:", number)      // 42
	fmt.Println("Pointer:", ptr)       // Address of 'number'
	fmt.Println("Dereferenced:", *ptr) // 42 (value at the memory address)

	// Modify the value using the pointer
	*ptr = 50
	fmt.Println("Updated Value:", number) // 50

	updateValue(&number)
	fmt.Println("Updating value by reference:", number)
}
