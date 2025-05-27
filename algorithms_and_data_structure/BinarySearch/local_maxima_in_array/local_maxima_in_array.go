package main

import (
	"fmt"
)

func localMaximaInArray(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 4, 2}
	index := localMaximaInArray(nums)
	fmt.Printf("Local maxima at index %d with value %d\n", index, nums[index])
}
