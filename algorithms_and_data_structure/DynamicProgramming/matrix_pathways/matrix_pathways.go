package main

import "fmt"

func matrixPathways(m int, n int) int {
	// Initialize the DP table with 1s
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = 1
		}
	}

	// Fill in the DP table
	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			dp[r][c] = dp[r-1][c] + dp[r][c-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	fmt.Println(matrixPathways(3, 7)) // Output: 28
}
