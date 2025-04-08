package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("GET method in GoLang!")

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error fetching the data", err)
		return
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	var post Post

	decoder.Decode(&post)

	// Print the post details,
	fmt.Println("Post details:")
	fmt.Println("ID:", post.ID)
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
	fmt.Println("UserID:", post.UserId)

}
