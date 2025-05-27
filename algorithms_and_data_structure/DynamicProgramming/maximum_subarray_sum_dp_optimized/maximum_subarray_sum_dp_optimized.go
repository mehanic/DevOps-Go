package main

import (
	"fmt"
)

func maximumSubarraySumDpOptimized(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	currentSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < n; i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumSubarraySumDpOptimized([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // Output: 6
}
