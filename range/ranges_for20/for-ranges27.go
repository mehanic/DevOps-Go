package main

import (
	"fmt"
	"math"
)

// Minimum function in Go with optional clipping
func minimum(clip *int, values ...int) int {
	if len(values) == 0 {
		return math.MaxInt // or any other default value
	}

	m := values[0]
	for _, value := range values[1:] {
		if value < m {
			m = value
		}
	}

	if clip != nil {
		if *clip > m {
			return *clip
		}
	}

	return m
}

func main() {
	// Calling minimum function without clipping
	fmt.Println(minimum(nil, 1, 5, 2, -5, 10))  // Returns -5

	// Calling minimum function with clipping
	clip := 0
	fmt.Println(minimum(&clip, 1, 5, 2, -5, 10)) // Returns 0
}

