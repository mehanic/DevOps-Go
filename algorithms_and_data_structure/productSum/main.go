package main

import "fmt"

// SpecialArray can hold both integers and nested arrays
type SpecialArray []interface{}

// O(n) time | O(d) space
// n - total elements including nested ones, d - maximum depth
func ProductSum(array SpecialArray) int {
	return helper(array, 1) // Start recursion with depth = 1
}

// Recursive helper function
func helper(array SpecialArray, multiplier int) int {
	sum := 0
	for _, el := range array {
		// If element is another SpecialArray (nested array)
		if cast, ok := el.(SpecialArray); ok {
			sum += helper(cast, multiplier+1) // Increase multiplier for deeper levels
		} else if cast, ok := el.(int); ok {
			sum += cast // Normal integer, just add it
		}
	}
	return sum * multiplier // Multiply sum by depth multiplier
}

func main() {
	// Example SpecialArray with nested structures
	example := SpecialArray{5, 2, SpecialArray{7, -1}, 3, SpecialArray{6, SpecialArray{4, -2}}}

	// Compute product sum
	result := ProductSum(example)

	// Print result
	fmt.Println("Product Sum:", result) // Expected Output: 41
}
