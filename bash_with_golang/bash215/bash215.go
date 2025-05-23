package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to print the board
func printBoard(board map[string]string) {
	fmt.Println(board["top-L"] + "|" + board["top-M"] + "|" + board["top-R"])
	fmt.Println("-+-+-")
	fmt.Println(board["mid-L"] + "|" + board["mid-M"] + "|" + board["mid-R"])
	fmt.Println("-+-+-")
	fmt.Println(board["low-L"] + "|" + board["low-M"] + "|" + board["low-R"])
}

func main() {
	// Initialize the board as a map
	theBoard := map[string]string{
		"top-L": " ", "top-M": " ", "top-R": " ",
		"mid-L": " ", "mid-M": " ", "mid-R": " ",
		"low-L": " ", "low-M": " ", "low-R": " ",
	}

	// Variables for the game
	turn := "X"
	reader := bufio.NewReader(os.Stdin)

	// Loop for 9 turns
	for i := 0; i < 9; i++ {
		printBoard(theBoard)
		fmt.Println("Turn for " + turn + ". Move on which space?")

		// Get the user's move
		move, _ := reader.ReadString('\n')
		move = strings.TrimSpace(move) // Remove newline character

		// Update the board with the player's move
		theBoard[move] = turn

		// Switch turns
		if turn == "X" {
			turn = "O"
		} else {
			turn = "X"
		}
	}

	// Print the final state of the board
	printBoard(theBoard)
}

