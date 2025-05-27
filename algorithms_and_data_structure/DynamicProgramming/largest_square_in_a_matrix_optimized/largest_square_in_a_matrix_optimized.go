package main

import (
	"fmt"
)

func largestSquareInMatrixOptimized(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	prevRow := make([]int, n)
	maxLen := 0

	for i := 0; i < m; i++ {
		currRow := make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				currRow[j] = matrix[i][j]
			} else if matrix[i][j] == 1 {
				currRow[j] = 1 + min(currRow[j-1], prevRow[j-1], prevRow[j])
			}
			if currRow[j] > maxLen {
				maxLen = currRow[j]
			}
		}
		prevRow = currRow
	}
	return maxLen * maxLen
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// Example usage
func main() {
	matrix := [][]int{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}
	fmt.Println(largestSquareInMatrixOptimized(matrix)) // Output: 9
}
