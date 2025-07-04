package main

import "fmt"

func findTheTargetInARotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] {
			// Left part is sorted
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right part is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	if len(nums) > 0 && nums[left] == target {
		return left
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(findTheTargetInARotatedSortedArray(nums, target)) // Output: 4
}
