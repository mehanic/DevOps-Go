package main

import "fmt"

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	res := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			res = max(res, dfs(r, c, matrix, memo))
		}
	}
	return res
}

func dfs(r, c int, matrix, memo [][]int) int {
	if memo[r][c] != 0 {
		return memo[r][c]
	}

	maxPath := 1
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		nextR, nextC := r+d[0], c+d[1]
		if isWithinBounds(nextR, nextC, matrix) && matrix[nextR][nextC] > matrix[r][c] {
			maxPath = max(maxPath, 1+dfs(nextR, nextC, matrix, memo))
		}
	}

	memo[r][c] = maxPath
	return maxPath
}

func isWithinBounds(r, c int, matrix [][]int) bool {
	return r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[0])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Optional: You can test the function in a main() if needed
func main() {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	fmt.Println(longestIncreasingPath(matrix)) // Output: 4
}
