package main

import (
	"fmt"
	"log"
)

// Function that performs division
func spam(divideBy float64) (float64, error) {
	if divideBy == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return 42 / divideBy, nil
}

func main() {
	// Test the spam function with different values
	for _, value := range []float64{2, 12, 0, 1} {
		result, err := spam(value)
		if err != nil {
			log.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
}

