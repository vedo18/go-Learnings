package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current time
	currentTime := time.Now()
	fmt.Println("Current time is: ", currentTime)

	// Format the current time using Go's reference layout
	// The format "2006-01-02 15:04:05" is Go's reference layout, used to define the desired time format.
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted time is: ", formattedTime)

	// Get the current date (Year, Month, Day)
	year, month, day := currentTime.Date()
	fmt.Printf("Current date is: %d-%d-%d\n", year, month, day)

	// Get the current time (Hour, Minute, Second)
	hour, minute, second := currentTime.Clock()
	fmt.Printf("Current time is: %02d:%02d:%02d\n", hour, minute, second)
}
