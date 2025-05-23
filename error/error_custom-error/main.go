package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	// Test cases
	testValues := []float64{16, 0, -10, 25, -7}

	for _, value := range testValues {
		result, err := sqrt(value)
		if err != nil {
			log.Printf("Error: %v (input: %v)\n", err, value)
			continue
		}
		fmt.Printf("The square root of %.2f is %.2f\n", value, result)
	}
}

// sqrt calculates the square root of a number, returning an error for negative inputs
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return math.Sqrt(f), nil
}


// The square root of 16.00 is 4.00
// The square root of 0.00 is 0.00
// 2025/02/14 15:06:38 Error: norgate math: square root of negative number (input: -10)
// The square root of 25.00 is 5.00
// 2025/02/14 15:06:38 Error: norgate math: square root of negative number (input: -7)
