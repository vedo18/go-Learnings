package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("Current time: ", currentTime)

	fmt.Println("Formatted time", currentTime.Format("01-02-2006 15:04:05 Monday")) //format will remain same
}
