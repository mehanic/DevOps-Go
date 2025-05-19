package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Define the options
	words := []string{"linux", "bash", "shell"}
	
	// Display the menu
	fmt.Println("Choose a word:")
	for i, word := range words {
		fmt.Printf("%d) %s\n", i+1, word) // Print the options as 1, 2, 3...
	}

	// Prompt for user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of your choice: ")
	input, _ := reader.ReadString('\n')

	// Convert input to integer
	choice, err := strconv.Atoi(input[:len(input)-1]) // Convert to int, ignore the newline character
	if err != nil || choice < 1 || choice > len(words) {
		fmt.Println("Invalid choice")
		return
	}

	// Display the selected word
	fmt.Printf("The word you have selected is: %s\n", words[choice-1])
}

