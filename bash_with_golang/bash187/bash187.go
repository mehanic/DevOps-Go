package main

import "fmt"

func main() {
	// Define the two arrays
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 6, 8}

	// Create a slice to store the zipped pairs
	var c [][]int

	// Loop through the arrays and create pairs
	for i := 0; i < len(a) && i < len(b); i++ {
		// Append the pairs of elements from a and b
		c = append(c, []int{a[i], b[i]})
	}

	// Print the zipped arrays
	fmt.Println(c)
}

