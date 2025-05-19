package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// choice function that mimics the Bash choice function
func choice(prompt string) string {
	fmt.Print(prompt)

	// Read input from the user
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')

	// Trim whitespace and newline characters
	answer = strings.TrimSpace(answer)

	// Standardize the response
	switch strings.ToLower(answer) {
	case "y", "1":
		return "y"
	case "n", "0":
		return "n"
	default:
		return answer // Return the raw input if it's not recognized
	}
}

func main() {
	// Example usage of the choice function
	result := choice("Do you want to play a game? (y/n): ")

	// Handle the result
	if result == "y" {
		fmt.Println("You chose to play the game!")
	} else if result == "n" {
		fmt.Println("You chose not to play the game.")
	} else {
		fmt.Printf("Unrecognized choice: %s\n", result)
	}
}

