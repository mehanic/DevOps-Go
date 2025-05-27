package main

import (
	"fmt"
	"math/rand"
	"time"
)

func kthLargestIntegerQuickselect(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickselect(nums, 0, len(nums)-1, k)
}

func quickselect(nums []int, left, right, k int) int {
	n := len(nums)
	if left >= right {
		return nums[left]
	}
	randomIndex := rand.Intn(right-left+1) + left
	nums[randomIndex], nums[right] = nums[right], nums[randomIndex]

	pivotIndex := partition(nums, left, right)

	if pivotIndex < n-k {
		return quickselect(nums, pivotIndex+1, right, k)
	} else if pivotIndex > n-k {
		return quickselect(nums, left, pivotIndex-1, k)
	} else {
		return nums[pivotIndex]
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	lo := left
	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[lo], nums[i] = nums[i], nums[lo]
			lo++
		}
	}
	nums[lo], nums[right] = nums[right], nums[lo]
	return lo
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(kthLargestIntegerQuickselect(nums, k)) // Output: 5
}
