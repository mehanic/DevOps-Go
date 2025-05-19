package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// choice function to read user input
func choice(prompt string) string {
	fmt.Print(prompt)

	// Read input from the user
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')

	// Trim whitespace and newline characters
	return strings.TrimSpace(answer)
}

func main() {
	// Prompt for favorite color
	color := choice("Enter your favorite color, if you have one: ")

	// Check if the input is not empty
	if color != "" {
		fmt.Printf("You chose: %s\n", color)
	} else {
		fmt.Println("You do not have a favorite color.")
	}
}

