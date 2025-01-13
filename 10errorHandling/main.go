package main

import "fmt"

func division(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("Denominator can't be zero ")
	}

	return a / b, nil
}

func main() {

	ans, err := division(10, 0)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Result", ans)

}
