package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	// Test cases
	fmt.Println(climbStairs(2)) // Output: 2
	fmt.Println(climbStairs(3)) // Output: 3
	fmt.Println(climbStairs(4)) // Output: 5
	fmt.Println(climbStairs(5)) // Output: 8
	fmt.Println(climbStairs(6)) // Output: 13
}
