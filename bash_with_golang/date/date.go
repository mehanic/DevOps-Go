package main

import (
	"fmt"
	"time"
)

func main() {
	// Get current time
	now := time.Now()

	// Unix epoch time in seconds
	fmt.Printf("Unix epoch time (seconds): %d\n", now.Unix())

	// Seconds of the current minute
	fmt.Printf("Current seconds (UTC): %02d\n", now.UTC().Second())

	// Full month name
	fmt.Printf("Month (full): %s\n", now.Format("January"))

	// Full day of the week name
	fmt.Printf("Day of the week (full): %s\n", now.Format("Monday"))

	// Short day of the week name
	fmt.Printf("Day of the week (short): %s\n", now.Format("Mon"))

	// Date in MM/DD/YY format
	fmt.Printf("Fixed date format (MM/DD/YY): %s\n", now.Format("01/02/06"))

	// Parsing specific date strings
	parsedTime1, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", "Thu Nov 18 08:07:21 IST 2010")
	if err != nil {
		fmt.Println("Error parsing date:", err)
	} else {
		// Print Unix epoch time of the parsed date
		fmt.Printf("Parsed date to Unix time (Thu Nov 18 08:07:21 IST 2010): %d\n", parsedTime1.Unix())
	}

	parsedTime2, err := time.Parse("Jan 2 2006", "Jan 20 2001")
	if err != nil {
		fmt.Println("Error parsing date:", err)
	} else {
		// Print the day of the week for the parsed date
		fmt.Printf("Parsed date (Jan 20 2001) Day of the week: %s\n", parsedTime2.Format("Monday"))
	}

	// Print today's date in "DD Month YYYY" format
	fmt.Printf("Today's date: %s\n", now.Format("02 January 2006"))
}

