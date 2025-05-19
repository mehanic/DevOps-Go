package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to convert a string to lowercase
func toLower(input string) string {
	return strings.ToLower(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt the user for input
		fmt.Print("Enter c to continue or q to exit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Remove newline and extra spaces

		// Convert input to lowercase
		input = toLower(input)

		// Check if input is "q" to exit the loop
		if input == "q" {
			break
		}
	}

	fmt.Println("Finished")
}

