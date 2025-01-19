package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "one, two, three, four, five"

	// Splitting the string into an array of strings

	arr := strings.Split(str, ", ")
	fmt.Println("Splitting the string into an array of strings", arr)

	//Counts the number of non-overlapping instances of a substring in a string.
	text := "banana"
	count := strings.Count(text, "a")
	fmt.Println("Count of 'a':", count) // 3

	//Joins elements of a slice into a single string with a specified separator.
	parts := []string{"hello", "world", "go"}
	result := strings.Join(parts, "-")
	fmt.Println("Joined Result:", result)

	//Converts a string to uppercase or lowercase.

	text1 := "Hello World"
	fmt.Println(strings.ToUpper(text1)) // "HELLO WORLD"
	fmt.Println(strings.ToLower(text1))

	//Removes leading and trailing characters (default is whitespace).
	text3 := "   hello world   "
	fmt.Println(strings.Trim(text3, " "))

}
