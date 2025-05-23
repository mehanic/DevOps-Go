package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current date and time
	now := time.Now()

	// Print the full date
	fmt.Printf("The date is %s\n", now.Format(time.RFC1123))

	// Extract and print the month
	month := now.Month()
	fmt.Printf("The month is %s\n", month.String())

	// Exit the program
}

