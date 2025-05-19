package main

import (
	"fmt"
	"os"
	"time"

	"github.com/manifoldco/promptui"
)

func main() {
	// Prompt for date selection
	prompt := promptui.Prompt{
		Label:     "Select a date (YYYY-MM-DD)",
		Templates: &promptui.PromptTemplates{},
	}

	// Run the prompt and handle user input
	date, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	// Parse the input date
	selectedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
		os.Exit(1)
	}

	// Print the selected date
	fmt.Printf("You selected: %s\n", selectedDate.Format("January 2, 2006"))

	// Return exit code based on success or failure
	if err == nil {
		os.Exit(0) // Success
	} else {
		os.Exit(1) // Failure
	}
}

