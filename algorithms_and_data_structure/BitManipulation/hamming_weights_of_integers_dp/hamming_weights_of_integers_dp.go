package main

import "fmt"

func hammingWeightsOfIntegersDP(n int) []int {
	dp := make([]int, n+1)
	for x := 1; x <= n; x++ {
		dp[x] = dp[x>>1] + (x & 1)
	}
	return dp
}

func main() {
	n := 5
	result := hammingWeightsOfIntegersDP(n)
	fmt.Println(result) // Output: [0 1 1 2 1 2]
}
