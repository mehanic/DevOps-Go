package main

import (
	"fmt"
)

func findTheInsertionIndex(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println("Insertion index:", findTheInsertionIndex(nums, target)) // Output: 2
}
