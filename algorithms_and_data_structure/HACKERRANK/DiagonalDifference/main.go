package main

import (
	"fmt"
)

func diagonalDifference(arr [][]int32) int32 {
	n := len(arr)
	var primary, secondary int32

	for i := 0; i < n; i++ {
		primary += arr[i][i]
		secondary += arr[i][n-1-i]
	}

	diff := primary - secondary
	if diff < 0 {
		return -diff
	}
	return diff
}

func main() {
	var n int
	fmt.Scan(&n)

	// First read into [][]int
	temp := make([][]int, n)
	for i := 0; i < n; i++ {
		temp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&temp[i][j])
		}
	}

	// Convert to [][]int32
	arr := make([][]int32, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int32, n)
		for j := 0; j < n; j++ {
			arr[i][j] = int32(temp[i][j])
		}
	}

	result := diagonalDifference(arr)
	fmt.Println(result)
}
