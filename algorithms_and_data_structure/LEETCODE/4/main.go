package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func main() {
	// Example slices
	example1 := []int{1, 2, 3, 4, 5}
	example2 := []int{1, 2, 3, 1}

	// Test case 1
	fmt.Println("Example 1:", example1)
	fmt.Println("Contains duplicate?", containsDuplicate(example1)) // false

	// Test case 2
	fmt.Println("Example 2:", example2)
	fmt.Println("Contains duplicate?", containsDuplicate(example2)) // true
}
