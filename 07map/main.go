package main

import "fmt"

func main() {
	numberToString := make(map[int]string)

	numberToString[0] = "zero"
	numberToString[1] = "one"
	numberToString[2] = "two"

	fmt.Println("Number to String: ", numberToString)

	for key, value := range numberToString {
		fmt.Printf("%d is %s\n", key, value)
		fmt.Printf("%s is %d\n", value, key)
	}
}
