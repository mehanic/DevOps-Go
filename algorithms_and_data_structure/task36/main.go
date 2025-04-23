package main

import (
	"fmt"
	"math"
)
//1. Recursion
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


//2. Dynamic Programming (Top-Down)
func minCostClimbingStairs(cost []int) int {
    memo := make([]int, len(cost))
    for i := 0; i < len(cost); i++ {
        memo[i] = -1
    }

    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= len(cost) {
            return 0
        }
        if memo[i] != -1 {
            return memo[i]
        }
        memo[i] = cost[i] + min(dfs(i+1), dfs(i+2))
        return memo[i]
    }
    
    return min(dfs(0), dfs(1))
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


//---

//3. Dynamic Programming (Bottom-Up)
func minCostClimbingStairs1(cost []int) int {
    n := len(cost)
    dp := make([]int, n+1)

    for i := 2; i <= n; i++ {
        dp[i] = min(dp[i-1] + cost[i-1],
                    dp[i-2] + cost[i-2])
    }

    return dp[n]
}

func min1(a, b int) int {
    if a < b {
        return a
    }
    return b
}

//4. Dynamic Programming (Space Optimized)

func minCostClimbingStairs2(cost []int) int {
    n := len(cost)
    for i := n - 3; i >= 0; i-- {
        cost[i] += min(cost[i+1], cost[i+2])
    }
    return min(cost[0], cost[1])
}

func min2(a, b int) int {
    if a < b {
        return a
    }
    return b
}
