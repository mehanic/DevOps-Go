package main

import "fmt"

func main() {
	// Define a tuple (using a slice in Go)
	dimensions := []int{200, 50}

	fmt.Println("Original dimensions:")
	for _, dimension := range dimensions {
		fmt.Println(dimension)
	}

	// Modify the dimensions (reassigning the slice)
	dimensions = []int{400, 100}
	
	fmt.Println("\nModified dimensions:")
	for _, dimension := range dimensions {
		fmt.Println(dimension)
	}
}

// Original dimensions:
// 200
// 50

// Modified dimensions:
// 400
// 100
