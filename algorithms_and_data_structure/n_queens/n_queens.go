package main

import "fmt"

var res int

func nQueens(n int) int {
	res = 0
	dfs(0, map[int]bool{}, map[int]bool{}, map[int]bool{}, n)
	return res
}

func dfs(r int, diagonalsSet, antiDiagonalsSet, colsSet map[int]bool, n int) {
	// Termination condition: we've placed all n queens
	if r == n {
		res++
		return
	}
	for c := 0; c < n; c++ {
		currDiagonal := r - c
		currAntiDiagonal := r + c
		// Skip if this column or diagonal is already occupied
		if colsSet[c] || diagonalsSet[currDiagonal] || antiDiagonalsSet[currAntiDiagonal] {
			continue
		}
		// Place queen
		colsSet[c] = true
		diagonalsSet[currDiagonal] = true
		antiDiagonalsSet[currAntiDiagonal] = true

		dfs(r+1, diagonalsSet, antiDiagonalsSet, colsSet, n)

		// Backtrack
		delete(colsSet, c)
		delete(diagonalsSet, currDiagonal)
		delete(antiDiagonalsSet, currAntiDiagonal)
	}
}

func main() {
	n := 8
	fmt.Printf("Number of solutions for %d-Queens: %d\n", n, nQueens(n))
}
