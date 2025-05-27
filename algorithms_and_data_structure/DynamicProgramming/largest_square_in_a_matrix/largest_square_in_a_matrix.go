package main

import (
	"fmt"
)

func largestSquareInMatrix(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	maxLen := 0

	// Base case: first row
	for j := 0; j < n; j++ {
		if matrix[0][j] == 1 {
			dp[0][j] = 1
			maxLen = 1
		}
	}

	// Base case: first column
	for i := 0; i < m; i++ {
		if matrix[i][0] == 1 {
			dp[i][0] = 1
			maxLen = 1
		}
	}

	// Fill the rest of the DP table
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen * maxLen
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// Example usage
func main() {
	matrix := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	fmt.Println(largestSquareInMatrix(matrix)) // Output: 9
}
