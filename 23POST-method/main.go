package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Post struct {
	// Defines a struct named Post to match the structure of the JSON response
	UserId int    `json:"userId"` // Field to store userId from JSON, the tag maps "userId" from JSON to UserId in Go
	Id     int    `json:"id"`     // Field to store id from JSON
	Title  string `json:"title"`  // Field to store title from JSON
	Body   string `json:"body"`   // Field to store body from JSON
}

func main() {
	fmt.Println("POST method in go lang!")

	post := Post{
		UserId: 18,
		Id:     7,
		Title:  "Testing in Progress",
		Body:   "Post method in go lang!",
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		fmt.Println("Error in json", err)
	}

	fmt.Println("json data", jsonData)

	jsonReader := strings.NewReader(string(jsonData))

	fmt.Println("json reader", jsonReader)

	url := "https://jsonplaceholder.typicode.com/posts/"

	res, err := http.Post(url, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error while doing post request", err)
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("json reader", data)

	fmt.Println("Response", string(data))

}
