package main

import (
	"fmt"
	"time"
)

func main() {
	main2()
}

func main1() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '#'
	)

	var cell rune // current cell (for caching)

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// create a drawing buffer
	buf := make([]rune, 0, width*height)

	// draw a smiley
	board[12][2] = true
	board[16][2] = true
	board[14][4] = true
	board[10][6] = true
	board[18][6] = true
	board[12][7] = true
	board[14][7] = true
	board[16][7] = true

	// use the loop for measuring the performance difference
	for i := 0; i < 100; i++ {
		// rewind the buffer so that the program reuses it
		buf = buf[:0]

		// draw the board into the buffer
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				// fmt.Print(string(cell), " ")
				buf = append(buf, cell, ' ')
			}
			// fmt.Println()
			buf = append(buf, '\n')
		}

		// print the buffer
		fmt.Print(string(buf))
	}
}

func main2() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '#'
	)

	var cell rune
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	buf := make([]rune, 0, width*height)

	for frame := 0; frame < 20; frame++ {
		// Очистим доску
		for x := range board {
			for y := range board[x] {
				board[x][y] = false
			}
		}

		// Анимируем "подмигивание"
		isWink := frame%2 == 0

		// Глаза
		board[12][2] = true
		if !isWink {
			board[16][2] = true // правый глаз открыт
		}

		// Нос
		board[14][4] = true

		// Щёки
		board[10][6] = true
		board[18][6] = true

		// Улыбка
		if isWink {
			// Поднятая улыбка
			board[12][7] = true
			board[13][6] = true
			board[14][6] = true
			board[15][6] = true
			board[16][7] = true
		} else {
			// Прямая линия
			board[12][7] = true
			board[14][7] = true
			board[16][7] = true
		}

		// Очистка буфера
		buf = buf[:0]

		// Отрисовка доски в буфер
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		// Очистка экрана (работает в большинстве терминалов)
		fmt.Print("\033[H\033[2J")

		// Печать
		fmt.Print(string(buf))

		time.Sleep(500 * time.Millisecond)
	}
}
