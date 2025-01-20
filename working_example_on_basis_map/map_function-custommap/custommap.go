package main

import (
	"fmt"
)

// Define a function type for the transformation
type Transformation func(int) int

// customMap applies a transformation function to each element of the slice
func customMap(transformation Transformation, array []int) []int {
	result := make([]int, len(array))
	for i, e := range array {
		result[i] = transformation(e)
	}
	return result
}

// addOne function increments the input integer by 1
func addOne(x int) int {
	return x + 1
}

// test function to test the customMap and addOne functions
func test() {
	myArray := []int{1, 2, 3, 4, 5}
	transformedArray := customMap(addOne, myArray)
	fmt.Println(transformedArray)
}

func main() {
	test()
}

// [2 3 4 5 6]
