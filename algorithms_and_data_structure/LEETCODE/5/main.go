package main

import "fmt"

// @leet start
func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}

// @leet end
func main() {
	// Test Case 1
	nums1 := []int{2, 2, 1}
	fmt.Println("Input:", nums1)
	fmt.Println("Single number:", singleNumber(nums1)) // Output: 1

	// Test Case 2
	nums2 := []int{4, 1, 2, 1, 2}
	fmt.Println("Input:", nums2)
	fmt.Println("Single number:", singleNumber(nums2)) // Output: 4

	// Test Case 3
	nums3 := []int{1}
	fmt.Println("Input:", nums3)
	fmt.Println("Single number:", singleNumber(nums3)) // Output: 1

	// Test Case 4
	nums4 := []int{0, 0, 5}
	fmt.Println("Input:", nums4)
	fmt.Println("Single number:", singleNumber(nums4)) // Output: 5

	// Test Case 5
	nums5 := []int{-1, -1, -2}
	fmt.Println("Input:", nums5)
	fmt.Println("Single number:", singleNumber(nums5)) // Output: -2
}
