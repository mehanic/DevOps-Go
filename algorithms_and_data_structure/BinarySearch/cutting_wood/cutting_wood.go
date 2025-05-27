package main

import (
	"fmt"
)

func cuttingWood(heights []int, k int) int {
	left, right := 0, max(heights)
	for left < right {
		mid := (left + right) / 2 + 1 // смещение в сторону правой границы
		if cutsEnoughWood(mid, k, heights) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return right
}

func cutsEnoughWood(H int, k int, heights []int) bool {
	woodCollected := 0
	for _, height := range heights {
		if height > H {
			woodCollected += height - H
		}
	}
	return woodCollected >= k
}

func max(nums []int) int {
	maxVal := nums[0]
	for _, n := range nums[1:] {
		if n > maxVal {
			maxVal = n
		}
	}
	return maxVal
}

func main() {
	heights := []int{20, 15, 10, 17}
	k := 7
	fmt.Println("Max cutting height:", cuttingWood(heights, k))
}
