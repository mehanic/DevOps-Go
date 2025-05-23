package main

import (
	"fmt"
)

// Define a function type for transformation
type Transformation func(int) int

// customMap applies a transformation function to each element of the slice
func customMap(transformation Transformation, array []int) []int {
	result := make([]int, len(array))
	for i, e := range array {
		result[i] = transformation(e)
	}
	return result
}

func main() {
	// Example transformation function: doubling the value
	double := func(x int) int {
		return x * 2
	}

	// Example slice
	numbers := []int{1, 2, 3, 4, 5}

	// Apply the transformation
	transformedNumbers := customMap(double, numbers)

	// Print the results
	fmt.Println(transformedNumbers)
}

