package main

import "fmt"

func knapsack(cap int, weights []int, values []int) int {
	n := len(values)
	// Initialize a 2D slice with zeros.
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, cap+1)
	}

	// Fill the DP table from the bottom up.
	for i := n - 1; i >= 0; i-- {
		for c := 1; c <= cap; c++ {
			if weights[i] <= c {
				include := values[i] + dp[i+1][c-weights[i]]
				exclude := dp[i+1][c]
				dp[i][c] = max(include, exclude)
			} else {
				dp[i][c] = dp[i+1][c]
			}
		}
	}
	return dp[0][cap]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage
func main() {
	weights := []int{1, 3, 4, 5}
	values := []int{1, 4, 5, 7}
	capacity := 7
	fmt.Println(knapsack(capacity, weights, values)) // Output: 9
}
