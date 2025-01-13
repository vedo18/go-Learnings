package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Give rating")

	rating, _ := reader.ReadString('\n')

	fmt.Println("Your rating", rating)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}
	updatedRating := numRating + 1

	fmt.Println("Your rating after adding 1 is:", updatedRating)

	// Check if the updated rating is 4.0 or above
	if updatedRating >= 4.0 {
		fmt.Println("Great, your rating is above or equal to 4.0")
	} else {
		fmt.Println("Not great, your rating is below 4.0")
	}

}
