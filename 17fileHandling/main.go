package main

import (
	"fmt"
	"os"
)

func main() {
	// Step 1: Create a file
	// os.Create() creates a file with the given name or truncates it if it already exists.
	// It returns a file descriptor and an error.
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function exits

	// Step 2: Write data to the file
	data := "Hello, Golang!\nThis is a sample file."
	_, err = file.WriteString(data) // WriteString writes the string to the file
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to file successfully!")

	// Step 3: Read data from the file
	// os.ReadFile() reads the entire file content into a byte slice.
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte slice to a string and print the file content
	fmt.Println("File content:")
	fmt.Println(string(content))
}
