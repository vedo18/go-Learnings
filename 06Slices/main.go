package main

import (
	"fmt"
)

func main() {

	sports := []string{"Football", "Cricket", "Tennis"}

	fmt.Println("Sports", sports)

	sports = append(sports, "Basketball")

	fmt.Println(sports)

	numbers := make([]int, 4)

	numbers[0] = 60
	numbers[1] = 20
	numbers[2] = 100
	numbers[3] = 40
	numbers[5] = 70

	fmt.Println(numbers)

	//sort
	// sort.Ints(numbers)
	// fmt.Println("After sorting", numbers)

	// fmt.Println(sort.IntsAreSorted(numbers))
	//delete
	index := 2

	numbers = append(numbers[:index], numbers[index+1:]...)

	fmt.Println("after delete", numbers)

}
