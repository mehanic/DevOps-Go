// Binary Search
// You are given an array of distinct integers nums, sorted in ascending order, and an integer target.

// Implement a function to search for target within nums. If it exists, then return its index, otherwise, return -1.

// Your solution must run in
// O
// (
// l
// o
// g
// n
// )
// O(logn) time.

// Example 1:

// Input: nums = [-1,0,2,4,6,8], target = 4

// Output: 3
// Example 2:

// Input: nums = [-1,0,2,4,6,8], target = 3

// Output: -1
// Constraints:

// 1 <= nums.length <= 10000.
// -10000 < nums[i], target < 10000

package main

import (
	"fmt"
)

// BinarySearch - O(log n) time | O(1) space
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2 // Avoids potential integer overflow

		if nums[mid] == target {
			return mid // Found target, return index
		} else if nums[mid] < target {
			left = mid + 1 // Search in the right half
		} else {
			right = mid - 1 // Search in the left half
		}
	}

	return -1 // Target not found
}

func main() {
	// Test cases
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{-1, 0, 2, 4, 6, 8}, 4},
		{[]int{-1, 0, 2, 4, 6, 8}, 3},
		{[]int{1, 3, 5, 7, 9}, 7},
		{[]int{-10, -5, 0, 5, 10}, -5},
		{[]int{10, 20, 30, 40}, 25},
	}

	// Running test cases
	for _, test := range testCases {
		fmt.Printf("Searching for %d in %v -> Index: %d\n",
			test.target, test.nums, BinarySearch(test.nums, test.target))
	}
}
