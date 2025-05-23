package main

import (
	"fmt"
//	"os"
	"time"
)

func main() {
	// Characters to be displayed
	chars := []rune{'-', '=', '|'}

	// Infinite loop
	for {
		for _, ch := range chars {
			// Move the cursor to row 33, column 0
			fmt.Printf("\033[33;0H") // ANSI escape code for cursor position (row 33, col 0)
			fmt.Print(string(ch))    // Print the character without a newline
			time.Sleep(1 * time.Second) // Sleep for 1 second
		}
	}
}

