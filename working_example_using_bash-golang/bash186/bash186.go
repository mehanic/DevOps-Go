package main

import (
	"fmt"
)

func main() {
	// Define the two arrays
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 6, 8}

	// Create a slice to store the pairs
	var c [][]int

	// Loop through the arrays and create pairs
	for idx := range a {
		// Append the pairs of elements from a and b to the slice
		c = append(c, []int{a[idx], b[idx]})
	}

	// Print the result
	fmt.Println(c)
}

