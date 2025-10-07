package main

import (
    "fmt"
)

func minCostClimbingStairs(cost []int) int {
    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= len(cost) {
            return 0
        }
        return cost[i] + min(dfs(i+1), dfs(i+2))
    }
    
    return min(dfs(0), dfs(1))
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    cost := []int{10, 15, 20}
    fmt.Println("Minimum cost to reach the top:", minCostClimbingStairs(cost))
}
