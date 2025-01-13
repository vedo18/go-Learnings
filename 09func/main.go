package main

import "fmt"

func addition(a, b int) int {
	return a + b
}

func main() {
	result := addition(10, 20)

	fmt.Println("result", result)
}
