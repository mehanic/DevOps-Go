package main

import (
	"fmt"
	"math"
)

func minCoinCombinationBottomUp(coins []int, target int) int {
	// Initialize DP array with "infinity" values.
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = math.MaxInt32
	}

	// Base case: 0 coins needed to make amount 0.
	dp[0] = 0

	// Fill the DP array
	for t := 1; t <= target; t++ {
		for _, coin := range coins {
			if coin <= t && dp[t-coin] != math.MaxInt32 {
				dp[t] = min(dp[t], 1+dp[t-coin])
			}
		}
	}

	if dp[target] == math.MaxInt32 {
		return -1
	}
	return dp[target]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{1, 2, 5}
	target := 11
	fmt.Println(minCoinCombinationBottomUp(coins, target)) // Output: 3 (5 + 5 + 1)
}
