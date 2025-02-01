package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Example URL
	rawURL := "https://example.com:8080/path/to/resource?name=John%20Doe&age=30#section1"

	// Step 1: Parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// Step 2: Access components of the URL
	fmt.Println("Scheme:", parsedURL.Scheme)     // https
	fmt.Println("Host:", parsedURL.Host)         // example.com:8080
	fmt.Println("Path:", parsedURL.Path)         // /path/to/resource
	fmt.Println("Fragment:", parsedURL.Fragment) // section1

}
