package main

import (
	"fmt"
)

func main() {
	// Define the options
	words := []string{"linux", "bash", "scripting", "tutorial"}

	// Display options
	fmt.Println("Choose one word:")
	for i, word := range words {
		fmt.Printf("%d) %s\n", i+1, word)
	}

	// Capture user input
	var choice int
	fmt.Print("Enter the number of your choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return
	}

	// Validate input and display selected word
	if choice < 1 || choice > len(words) {
		fmt.Println("Invalid choice. Please select a number between 1 and", len(words))
	} else {
		selectedWord := words[choice-1] // Adjust for zero-based index
		fmt.Printf("The word you have selected is: %s\n", selectedWord)
	}
}

