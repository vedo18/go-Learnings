package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15, 20}

	for index, value := range numbers {
		fmt.Printf("Index is %d and value is %d\n", index, value)
	}
}

// In go, only for loop is used
//in Go's for range loop, the first variable is always the index (or key for maps),
//and the second variable is always the value.
