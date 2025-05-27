package main

import "fmt"

// @leet start
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] < 10 {
			return digits
		}
		digits[i] = 0
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func main() {
	examples := [][]int{
		{1, 2, 3},    // 123 + 1 = 124
		{4, 3, 2, 1}, // 4321 + 1 = 4322
		{9},          // 9 + 1 = 10
		{9, 9, 9},    // 999 + 1 = 1000
	}

	for _, digits := range examples {
		fmt.Printf("Input: %v, Output: %v\n", digits, plusOne(digits))
	}
}
