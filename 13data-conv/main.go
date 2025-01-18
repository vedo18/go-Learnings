package main

import (
	"fmt"
	"strconv"
)

func main() {
	//int to float64
	var num int = 40
	fmt.Println("Value is", num)
	fmt.Printf("Type is %T\n", num)

	var data float64 = float64(num)
	data = data + 1.31
	fmt.Println("Value is", data)
	fmt.Printf("Type is %T\n", data)

	//int to string
	num1 := 40
	fmt.Println("Value is", num1)
	fmt.Printf("Type is %T\n", num1)

	data1 := strconv.Itoa(num1)
	fmt.Println("Value is", data1)
	fmt.Printf("Type is %T\n", data1)

	//string to int
	data2 := "123456789"
	fmt.Println("Value is", data2)
	fmt.Printf("Type is %T\n", data2)

	num2, _ := strconv.Atoi(data2)
	fmt.Println("Value is", num2)
	fmt.Printf("Type is %T\n", num2)

}
