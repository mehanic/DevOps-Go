package main

import "fmt"

// @leet start
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	blocks := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {

		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		blocks[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {

			val := board[r][c]
			if val == '.' {
				continue
			}
			blockId := getBlockNumber(r, c)
			if rows[r][val] || cols[c][val] || blocks[blockId][val] {
				return false
			}

			rows[r][val] = true
			cols[c][val] = true
			blocks[blockId][val] = true

		}
	}

	return true
}

func getBlockNumber(row, col int) int {
	return (row/3)*3 + col/3
}

func main() {

	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	valid := isValidSudoku(board)
	fmt.Println("Valid Sudoku:", valid)

	fmt.Println()
	main1()

	// TODO: implement
}

// @leet end
func main1() {

	board := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	valid := isValidSudoku(board)
	fmt.Println("Valid Sudoku:", valid)

	// TODO: implement
}
