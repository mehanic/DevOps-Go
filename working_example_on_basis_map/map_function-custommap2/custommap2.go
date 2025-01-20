package main

import (
	"fmt"
)

// multiply function that takes two integers and returns their product
func multiply(x, y int) int {
	return x * y
}

// customMap multiplies corresponding elements from two slices and returns a new slice
func customMap(xs, ys []int) []int {
	// Determine the length of the smaller slice to avoid out-of-range errors
	length := len(xs)
	if len(ys) < length {
		length = len(ys)
	}

	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = multiply(xs[i], ys[i])
	}
	return result
}

// test function to test the customMap and lambda functionality
func test() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ys := []int{2, 3, 4, 5, 6, 7}
	resultList := customMap(xs, ys)
	fmt.Println(resultList)
}

func main() {
	test()
}

// [2 6 12 20 30 42]
