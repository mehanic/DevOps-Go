package main

import (
	"fmt"
)

func firstAndLastOccurrences(nums []int, target int) []int {
	lower := lowerBoundBinarySearch(nums, target)
	upper := upperBoundBinarySearch(nums, target)
	return []int{lower, upper}
}

func lowerBoundBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if len(nums) > 0 && nums[left] == target {
		return left
	}
	return -1
}

func upperBoundBinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2 + 1 // bias to the right
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			left = mid
		}
	}
	if len(nums) > 0 && nums[right] == target {
		return right
	}
	return -1
}

func main() {
	nums := []int{1, 2, 2, 2, 3, 4, 5}
	target := 2
	fmt.Println(firstAndLastOccurrences(nums, target)) // Output: [1 3]
}
