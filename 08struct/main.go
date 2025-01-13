package main

import "fmt"

type User struct {
	Name     string
	Email    string
	Age      int
	isActive bool
}

func main() {

	saket := User{"Saket", "saketvats18@gmail.com", 24, true}

	fmt.Printf("Saket details are: %+v\n", saket)

}
