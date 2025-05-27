package main

import "fmt"

var memo = map[int]int{}

func climbingStairsTopDown(n int) int {
	// Base cases
	if n <= 2 {
		return n
	}
	// Check memoization map
	if val, exists := memo[n]; exists {
		return val
	}
	// Recursive calls with memoization
	memo[n] = climbingStairsTopDown(n-1) + climbingStairsTopDown(n-2)
	return memo[n]
}

func main() {
	fmt.Println(climbingStairsTopDown(5)) // Output: 8
}
