package main

import "fmt"

func main() {
	var num int = 40
	fmt.Println("Value is", num)
	fmt.Printf("Type is %T\n", num)

	var data float64 = float64(num)
	data = data + 1.31
	fmt.Println("Value is", data)
	fmt.Printf("Type is %T\n", data)
}
