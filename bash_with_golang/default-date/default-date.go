package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	// Calculate the starting date (noon last Monday)
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 { // Sunday
		weekday = 7
	}
	startDate := now.AddDate(0, 0, -1*weekday+1).Truncate(24 * time.Hour).Add(12 * time.Hour)

	startDateStr := startDate.Format("2006-01-02")

	// Regular expression to validate date format YYYY-MM-DD
	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

	var answer string
	for {
		fmt.Printf("The starting date is %s, is that correct? (Y/new date): ", startDateStr)
		_, err := fmt.Scanln(&answer)
		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		// Validate the answer
		switch {
		case answer == "Y" || answer == "y":
			break
		case dateRegex.MatchString(answer):
			fmt.Printf("Overriding %s with %s\n", startDateStr, answer)
			startDateStr = answer
		default:
			fmt.Println("Invalid date, please try again...")
			continue
		}
		break
	}

	// Calculate the end date (7 days after the start date)
	startDateParsed, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		os.Exit(1)
	}

	endDate := startDateParsed.AddDate(0, 0, 7)

	// Print the start and end dates
	fmt.Printf("START_DATE: %s\n", startDateStr)
	fmt.Printf("END_DATE:   %s\n", endDate.Format("2006-01-02"))
}

