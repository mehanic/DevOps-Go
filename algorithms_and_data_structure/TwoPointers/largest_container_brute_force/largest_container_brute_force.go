package main

import (
	"fmt"
)

// largestContainerBruteForce calculates the maximum water container area (brute force)
func largestContainerBruteForce(heights []int) int {
	n := len(heights)
	maxWater := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			height := heights[i]
			if heights[j] < height {
				height = heights[j]
			}
			width := j - i
			water := height * width

			if water > maxWater {
				maxWater = water
			}
		}
	}

	return maxWater
}

func main() {
	heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("Max area:", largestContainerBruteForce(heights))
}
