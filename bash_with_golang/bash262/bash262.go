package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Prompt the user for their first and last names
	fmt.Print("Please enter your first name and last name: ")

	// Read input from the user
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim the newline character at the end of input
	input = strings.TrimSpace(input)

	// Split the input into first name and last name
	names := strings.SplitN(input, " ", 2)
	if len(names) < 2 {
		fmt.Println("Please provide both first name and last name.")
		return
	}

	firstname := names[0]
	lastname := names[1]

	// Print the greeting
	fmt.Printf("Hello, %s. How is the %s family?\n", firstname, lastname)
}

