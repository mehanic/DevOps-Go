package main

import (
	"fmt"
	"math"
)

func maximumSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := math.Inf(-1)
	currentSum := math.Inf(-1)

	for _, num := range nums {
		currentSum = math.Max(currentSum+float64(num), float64(num))
		maxSum = math.Max(maxSum, currentSum)
	}

	return int(maxSum)
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maximumSubarraySum(arr)) // Output: 6
}
