package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for day
	fmt.Print("Enter day (1-31): ")
	dayStr, _ := reader.ReadString('\n')
	dayStr = strings.TrimSpace(dayStr)
	day, err := strconv.Atoi(dayStr)
	if err != nil || day < 1 || day > 31 {
		fmt.Println("Invalid day")
		return
	}

	// Prompt for month
	fmt.Print("Enter month (1-12): ")
	monthStr, _ := reader.ReadString('\n')
	monthStr = strings.TrimSpace(monthStr)
	month, err := strconv.Atoi(monthStr)
	if err != nil || month < 1 || month > 12 {
		fmt.Println("Invalid month")
		return
	}

	// Prompt for year
	fmt.Print("Enter year (e.g., 2018): ")
	yearStr, _ := reader.ReadString('\n')
	yearStr = strings.TrimSpace(yearStr)
	year, err := strconv.Atoi(yearStr)
	if err != nil || year < 1 {
		fmt.Println("Invalid year")
		return
	}

	// Display the selected date
	selectedDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	fmt.Printf("You selected: %s\n", selectedDate.Format("January 2, 2006"))
}

