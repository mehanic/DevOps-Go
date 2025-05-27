package main

import (
	"fmt"
)

func longestPalindromeInAString(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	// Initialize 2D DP table with false
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	maxLen := 1
	startIndex := 0

	// Base case: every single character is a palindrome
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// Base case: check for substrings of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			startIndex = i
			maxLen = 2
		}
	}

	// Check substrings of length 3 or more
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				startIndex = i
				maxLen = length
			}
		}
	}

	return s[startIndex : startIndex+maxLen]
}

func main() {
	s := "babad"
	fmt.Println(longestPalindromeInAString(s)) // Output: "bab" or "aba"
}
