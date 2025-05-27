package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Brute force approach
func largestContainerBruteForce(heights []int) int {
	n := len(heights)
	maxWater := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			height := min(heights[i], heights[j])
			width := j - i
			area := height * width
			if area > maxWater {
				maxWater = area
			}
		}
	}

	return maxWater
}

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Max water (brute force):", largestContainerBruteForce(heights))
}
