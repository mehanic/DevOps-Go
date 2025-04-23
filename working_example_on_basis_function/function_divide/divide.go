package main

import (
	"errors"
	"fmt"
)

func spam(divideBy int) (float64, error) {
	if divideBy == 0 {
		return 0, errors.New("division by zero")
	}
	return 42 / float64(divideBy), nil
}

func main() {
	// Define a slice of divisors to test with.
	divisors := []int{2, 12, 0, 1}

	for _, d := range divisors {
		result, err := spam(d)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
}

