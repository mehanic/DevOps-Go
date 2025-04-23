package main

import "fmt"

// Function to calculate the number of ways to climb to the top of the staircase
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	// Initialize the base cases
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2

	// Fill the dp array based on the recurrence relation
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
