package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Prompt for the user's name
	fmt.Print("What is your name? ")

	// Read the input from the user
	input, _ := reader.ReadString('\n')

	// Trim any whitespace or newline characters
	input = strings.TrimSpace(input)

	// Split the input into first name, middle name, and last name
	names := strings.Split(input, " ")

	// Assign names to variables (ensure there are enough names provided)
	var fname, mname, lname string
	fname = names[0] // First name
	if len(names) > 2 {
		mname = names[1] // Middle name
		lname = names[2] // Last name
	} else if len(names) == 2 {
		lname = names[1] // Last name only if no middle name
	}

	// Output the names
	fmt.Printf("Your first name is : %s\n", fname)
	fmt.Printf("Your middle name is : %s\n", mname)
	fmt.Printf("Your last name is : %s\n", lname)
}

