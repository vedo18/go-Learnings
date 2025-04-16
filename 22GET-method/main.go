package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("Get method in go lang!")

	url := "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Error while fetching", err)
		return
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	var post Post

	decoder.Decode(&post)

	fmt.Println("User id:", post.UserId)
	fmt.Println("Id:", post.Id)
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)

}
