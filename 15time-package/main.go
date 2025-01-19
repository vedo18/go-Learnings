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

	// Add 1 day to the current time
	tomorrow := currentTime.Add(24 * time.Hour)
	fmt.Println("Tomorrow's date and time: ", tomorrow)

	// Subtract 1 day from the current time
	yesterday := currentTime.Add(-24 * time.Hour)
	fmt.Println("Yesterday's date and time: ", yesterday)

	// Parse a time string into a Time object
	timeString := "2025-01-15 10:30:00"
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	} else {
		fmt.Println("Parsed time is: ", parsedTime)
	}

	// Sleep for 2 seconds
	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)
	fmt.Println("Awake after sleeping for 2 seconds!")

	// Get the Unix timestamp
	unixTimestamp := currentTime.Unix()
	fmt.Printf("Unix timestamp: %d\n", unixTimestamp)

	// Convert Unix timestamp back to Time
	timeFromUnix := time.Unix(unixTimestamp, 0)
	fmt.Println("Time from Unix timestamp: ", timeFromUnix)
}
