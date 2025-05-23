package main

import "fmt"

// add function that accepts a variadic parameter of integers
func add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	// Call the add function with different inputs
	fmt.Println(add(1, 2, 3))    // Output: 6
	fmt.Println(add(1, 2, 3, 4)) // Output: 10

	nums := []int{5, 6, 7, 2, 3}
	fmt.Println(add(nums...)) // Output: 23
}

// 6
// 10
// 23
