package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Step 1: Define the URL of the API
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Step 2: Make the GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer response.Body.Close() // Close the response body when done

	// Step 3: Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Step 4: Convert the response body to a string and print it
	fmt.Println("Response:", string(body))
}
