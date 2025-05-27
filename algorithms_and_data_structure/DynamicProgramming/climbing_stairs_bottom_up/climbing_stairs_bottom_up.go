package main

import "fmt"

func climbingStairsBottomUp(n int) int {
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	// Base cases
	dp[1], dp[2] = 1, 2

	// Fill the dp array from step 3 to n
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	fmt.Println(climbingStairsBottomUp(5)) // Output: 8
}
