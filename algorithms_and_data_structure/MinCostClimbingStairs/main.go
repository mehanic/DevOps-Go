package main

import (
	"fmt"
	"math"
)

// MinCostClimbingStairs calculates the minimum cost to reach the top
func MinCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 2 {
		return int(math.Min(float64(cost[0]), float64(cost[1])))
	}

	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + int(math.Min(float64(dp[i-1]), float64(dp[i-2])))
	}

	return int(math.Min(float64(dp[n-1]), float64(dp[n-2])))
}

func main() {
	// Example 1
	cost1 := []int{1, 2, 3}
	fmt.Println("Minimum cost:", MinCostClimbingStairs(cost1)) // Output: 2

	// Example 2
	cost2 := []int{1, 2, 1, 2, 1, 1, 1}
	fmt.Println("Minimum cost:", MinCostClimbingStairs(cost2)) // Output: 4

	// Additional Example
	cost3 := []int{10, 15, 20}
	fmt.Println("Minimum cost:", MinCostClimbingStairs(cost3)) // Output: 15

	cost4 := []int{0, 1, 2, 2}
	fmt.Println("Minimum cost:", MinCostClimbingStairs(cost4)) // Output: 2
}
