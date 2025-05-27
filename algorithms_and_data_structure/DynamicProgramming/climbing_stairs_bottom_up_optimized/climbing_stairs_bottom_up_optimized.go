package main

import "fmt"

func climbingStairsBottomUpOptimized(n int) int {
	if n <= 2 {
		return n
	}

	oneStepBefore, twoStepsBefore := 2, 1
	var current int
	for i := 3; i <= n; i++ {
		current = oneStepBefore + twoStepsBefore
		twoStepsBefore = oneStepBefore
		oneStepBefore = current
	}

	return oneStepBefore
}

func main() {
	fmt.Println(climbingStairsBottomUpOptimized(5)) // Output: 8
}
