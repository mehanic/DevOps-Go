package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Do you like this question? ")
	replyVariable, _ := reader.ReadString('\n')
	replyVariable = strings.TrimSpace(replyVariable) // Remove any trailing newline or spaces

	// Convert the reply to lowercase for comparison
	replyLower := strings.ToLower(replyVariable)

	// Check if the user responded positively
	if replyLower == "y" || replyLower == "yes" {
		fmt.Println("Great, I worked really hard on it!")
		os.Exit(0)
	}

	// Check if the user responded negatively
	if replyLower == "n" || replyLower == "no" {
		fmt.Println("You did not? But I worked so hard on it!")
		os.Exit(0)
	}

	// If we get here, the user did not give a proper response
	fmt.Println("Please use yes/no!")
	os.Exit(1)
}

