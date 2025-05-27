package main

import "fmt"

func countIslands(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	count := 0
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == 1 {
				dfs(r, c, matrix)
				count++
			}
		}
	}
	return count
}

func dfs(r, c int, matrix [][]int) {
	matrix[r][c] = -1 // mark as visited
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, d := range dirs {
		nr, nc := r+d[0], c+d[1]
		if isWithinBounds(nr, nc, matrix) && matrix[nr][nc] == 1 {
			dfs(nr, nc, matrix)
		}
	}
}

func isWithinBounds(r, c int, matrix [][]int) bool {
	return r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[0])
}

// Example usage
func main() {
	matrix := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	fmt.Println(countIslands(matrix)) // Output: 3
}
