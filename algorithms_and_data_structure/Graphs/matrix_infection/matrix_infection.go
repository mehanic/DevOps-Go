package main

import (
	"fmt"
)

func matrixInfection(matrix [][]int) int {
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][2]int{}
	ones := 0
	seconds := 0

	rows := len(matrix)
	if rows == 0 {
		return -1
	}
	cols := len(matrix[0])

	// Count initial uninfected cells and enqueue infected ones
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 1 {
				ones++
			} else if matrix[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
		}
	}

	// BFS traversal
	for len(queue) > 0 && ones > 0 {
		seconds++
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			cell := queue[0]
			queue = queue[1:]
			r, c := cell[0], cell[1]

			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if isWithinBounds(nr, nc, matrix) && matrix[nr][nc] == 1 {
					matrix[nr][nc] = 2
					ones--
					queue = append(queue, [2]int{nr, nc})
				}
			}
		}
	}

	if ones == 0 {
		return seconds
	}
	return -1
}

func isWithinBounds(r, c int, matrix [][]int) bool {
	return r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[0])
}

func main() {
	matrix := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}
	fmt.Println(matrixInfection(matrix)) // Output: 4
}
