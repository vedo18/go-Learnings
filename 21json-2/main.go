package main

import (
    "encoding/json"
    "fmt"
)

// Define a struct that matches the structure of the JSON we want to parse
// The json tags define the mapping between JSON field names and struct field names
type User struct {
    Name  string `json:"name"`   // Maps to "name" in JSON
    Age   int    `json:"age"`    // Maps to "age" in JSON
    Email string `json:"email"`  // Maps to "email" in JSON
}

func main() {
    fmt.Println("JSON Unmarshalling Example")
    
    // This is our JSON data as a string
    // In a real application, this might come from a file, API response, etc.
    jsonString := `{"name":"Saket Vats","age":24,"email":"saketvats18@gmail.com"}`
    
    // Convert the JSON string to a byte slice, which is what Unmarshal requires
    jsonBytes := []byte(jsonString)
    
    // Create an empty User struct that will be filled with the JSON data
    var user User
    
    // Unmarshal parses the JSON and fills the struct with the data
    // We pass a pointer to our user variable (&user) so that Unmarshal can modify it
    err := json.Unmarshal(jsonBytes, &user)
    
    // Check if an error occurred during unmarshalling
    if err != nil {
        fmt.Println("Error during unmarshalling:", err)
    } else {
        // Print the resulting struct fields to confirm the data was parsed correctly
        fmt.Println("Parsed User data:")
        fmt.Println("Name:", user.Name)    // Should print: Name: Saket Vats
        fmt.Println("Age:", user.Age)      // Should print: Age: 24
        fmt.Println("Email:", user.Email)  // Should print: Email: saketvats18@gmail.com
        
        // We can now use the user struct in our Go program
        fmt.Printf("%s is %d years old\n", user.Name, user.Age)
    }
}