package main

import (
	"fmt"
	"os"
	"time"
)

// Function to display the progress bar
func progress(percent, max int) {
	column := 71 * percent / max
	nspace := 71 - column

	// Start the progress bar string
	bar := "\r["

	// Append '=' for completed part
	for i := 0; i < column; i++ {
		bar += "="
	}

	// Append '>' for the current position
	bar += ">"

	// Append spaces for the remaining part
	for i := 0; i < nspace; i++ {
		bar += " "
	}

	// Append the closing bracket and percentage
	bar += fmt.Sprintf("] %d/%d", percent, max)

	// Print the progress bar
	fmt.Print(bar)
	// Force the output to flush
	os.Stdout.Sync()
}

func main() {
	MAX := 100 // Maximum value

	for i := 0; i <= MAX; i++ {
		progress(i, MAX)
		time.Sleep(50 * time.Millisecond) // Simulate work
	}
	fmt.Println() // Move to the next line after completion

	// Exit program
}

