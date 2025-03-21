package main

import (
    "encoding/json"  // Package for JSON encoding and decoding
    "fmt"           
)

// Define a struct called User with three fields
// The `json:"..."` tags tell the JSON encoder what field names to use in the JSON output
type User struct {
    Name  string `json:"name"`   // Will appear as "name" in JSON
    Age   int    `json:"age"`    // Will appear as "age" in JSON
    Email string `json:"email"`  // Will appear as "email" in JSON
}

func main() {
    fmt.Println("JSON in go lang!")  // Print a simple message
    
    // Create a new User instance with sample data
    user := User{
        Name:  "Saket Vats", 
        Age:   24, 
        Email: "saketvats18@gmail.com",
    }
    
    // Marshal converts the Go struct into a JSON byte slice
    // jsonData will contain the JSON as bytes
    // err will contain any error that occurred during marshalling
    jsonData, err := json.Marshal(user)
    
    // Check if an error occurred during marshalling
    if err != nil {
        fmt.Println("Error:", err)  // Print the error if one occurred
    } else {
        // If no error, print the JSON data
        // string(jsonData) converts the byte slice to a readable string
        fmt.Println("JSON data:", string(jsonData))
        // This will output: JSON data: {"name":"Saket Vats","age":24,"email":"saketvats18@gmail.com"}
    }
}