package main

import (
	"fmt"
	"os"
)

func main() {
	// Constants for time conversions
	const (
		secsInDay   = 86400
		secsInHour  = 3600
		minsInHour  = 60
		secsInMin   = 60
	)

	// Check if the number of arguments is correct
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run time_converter.go <seconds>")
		os.Exit(1)
	}

	// Convert the first argument to an integer
	var totalSeconds int
	_, err := fmt.Sscanf(os.Args[1], "%d", &totalSeconds)
	if err != nil {
		fmt.Println("Error: Invalid input. Please provide a valid number of seconds.")
		os.Exit(1)
	}

	// Calculate days, hours, minutes, and seconds
	days := totalSeconds / secsInDay
	secs := totalSeconds % secsInDay
	hours := secs / secsInHour
	secs %= secsInHour
	minutes := secs / secsInMin
	seconds := secs % secsInMin

	// Print the result in the format D:H:M:S
	fmt.Printf("%d:%02d:%02d:%02d\n", days, hours, minutes, seconds)
}

