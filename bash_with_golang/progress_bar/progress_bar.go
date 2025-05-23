package main

import (
	"fmt"
	"time"
)

// progress displays a progress bar
func progress(percent, gt int) {
	column := (71 * percent) / gt
	nspace := 71 - column

	bar := "\r["
	for i := 0; i < column; i++ {
		bar += "="
	}
	bar += ">"

	for i := 0; i < nspace; i++ {
		bar += " "
	}
	bar += fmt.Sprintf("] %d/%d", percent, gt)

	// Print the progress bar
	fmt.Print(bar)
}

func main() {
	const MAX = 100 // Maximum value

	for i := 0; i <= MAX; i++ {
		progress(i, MAX) // Update the progress bar
		time.Sleep(50 * time.Millisecond) // Simulate work being done
	}
	fmt.Println() // Print a new line after progress bar completion
}

