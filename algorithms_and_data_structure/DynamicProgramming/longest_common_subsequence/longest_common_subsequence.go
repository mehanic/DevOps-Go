package main

import (
	"fmt"
)

func longestCommonSubsequenceOptimized(s1, s2 string) int {
	m, n := len(s1), len(s2)
	prevRow := make([]int, n+1)

	for i := m - 1; i >= 0; i-- {
		currRow := make([]int, n+1)
		for j := n - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				currRow[j] = 1 + prevRow[j+1]
			} else {
				currRow[j] = max(prevRow[j], currRow[j+1])
			}
		}
		prevRow = currRow
	}
	return prevRow[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Example usage
func main() {
	s1 := "abcde"
	s2 := "ace"
	fmt.Println(longestCommonSubsequenceOptimized(s1, s2)) // Output: 3
}
