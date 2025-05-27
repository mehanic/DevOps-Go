package main

import "fmt"

func matrixPathwaysOptimized(m int, n int) int {
	// Initialize prevRow as all 1s
	prevRow := make([]int, n)
	for i := 0; i < n; i++ {
		prevRow[i] = 1
	}

	// Iterate through the matrix starting from row 1
	for r := 1; r < m; r++ {
		currRow := make([]int, n)
		currRow[0] = 1
		for c := 1; c < n; c++ {
			currRow[c] = prevRow[c] + currRow[c-1]
		}
		prevRow = currRow
	}

	// Return the value in the bottom-right corner
	return prevRow[n-1]
}

func main() {
	fmt.Println(matrixPathwaysOptimized(3, 7)) // Output: 28
}
