package main

import (
	"fmt"
)

// 1. Brute Force
// Brute-force approach to find two numbers that sum to target
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j} // Return indices of the two numbers
			}
		}
	}
	return []int{} // Return empty if no solution
}

func main() {
	// Test cases
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))   // [0, 1] -> 2 + 7 = 9
	fmt.Println(twoSum([]int{3, 2, 4}, 6))        // [1, 2] -> 2 + 4 = 6
	fmt.Println(twoSum([]int{1, 5, 3, 8}, 9))     // [0, 3] -> 1 + 8 = 9
	fmt.Println(twoSum([]int{3, 3}, 6))           // [0, 1] -> 3 + 3 = 6
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 10)) // [] (no solution)
	fmt.Println("-----")
}
