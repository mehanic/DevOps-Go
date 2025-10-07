package main

import (
	"fmt"
)

func climbStairs(n int) int {
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= n {
			if i == n {
				return 1
			}
			return 0
		}
		return dfs(i+1) + dfs(i+2)
	}
	return dfs(0)
}

func main() {
	n := 4
	result := climbStairs(n)
	fmt.Printf("Number of ways to climb %d stairs: %d\n", n, result)
}
