package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current day of the week
	today := time.Now().Weekday()

	// Print the day in a sentence
	fmt.Printf("Today is %s\n", today)
}

