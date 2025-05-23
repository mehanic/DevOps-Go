package main

import (
	"fmt"
	"time"
)

// Function to get yesterday's date
func yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1) // Subtract 1 day from the current date
}

func main() {
	// Call the yesterday function and print the result
	fmt.Println("Yesterday's date:", yesterday().Format("2006-01-02")) // Format to YYYY-MM-DD
}

