package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	// Defines a struct named Post to match the structure of the JSON response
	UserId int    `json:"userId"` // Field to store userId from JSON, the tag maps "userId" from JSON to UserId in Go
	Id     int    `json:"id"`     // Field to store id from JSON
	Title  string `json:"title"`  // Field to store title from JSON
	Body   string `json:"body"`   // Field to store body from JSON
}

func main() {
	fmt.Println("Get method in go lang!")

	url := "https://jsonplaceholder.typicode.com/posts/1" // Defines the URL to make the request to

	response, err := http.Get(url) // Makes an HTTP GET request to the URL and returns response and error
	if err != nil {
		fmt.Println("Error while fetching", err)
		return
	}

	defer response.Body.Close() // Schedules closing of the response body when the function exits

	decoder := json.NewDecoder(response.Body) // Creates a new JSON decoder that reads from the response body

	var post Post         // Declares a variable of type Post to store the decoded JSON
	decoder.Decode(&post) // Decodes the JSON from the response body into the post variable

	fmt.Println("User id:", post.UserId) // Prints the userId field from the decoded JSON
	fmt.Println("Id:", post.Id)          // Prints the id field from the decoded JSON
	fmt.Println("Title:", post.Title)    // Prints the title field from the decoded JSON
	fmt.Println("Body:", post.Body)      // Prints the body field from the decoded JSON
}
