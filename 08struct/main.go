package main

import (
	"fmt"
)

type User struct {
	Name     string
	Age      int
	IsActive bool
}

type Contact struct {
	Email       string
	PhoneNumber string
}

type Address struct {
	Street  string
	City    string
	State   string
	Pincode int
}

type FullDetails struct {
	User
	Contact
	Address
}

func main() {

	var details FullDetails

	details.User = User{
		Name:     "Saket",
		Age:      24,
		IsActive: true,
	}

	details.Contact = Contact{
		Email:       "saketvats18@gmail.com",
		PhoneNumber: "1234567890",
	}

	details.Address = Address{
		Street:  "123 Main Street",
		City:    "Bangalore",
		State:   "Karnataka",
		Pincode: 560100,
	}

	fmt.Printf("Full Details: %+v\n", details)

	fmt.Println("Name:", details.Name)
	fmt.Println("Email:", details.Email)
	fmt.Println("City:", details.City)

}

// other ways to create struct

// Creates a pointer to a User struct with zero values
// user := new(User) Creates a pointer to a User struct with zero values
// user.Name = "Saket"
// user.Age = 24
// user.IsActive = true

// user := User{Name: "Saket", Age: 24, IsActive: true}

// details := FullDetails{
// 	User: User{Name: "Saket", Age: 24, IsActive: true},
// }
