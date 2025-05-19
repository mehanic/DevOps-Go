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

	// Prompt and read the first name
	fmt.Print("Enter first name: ")
	firstName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName) // Remove any trailing newline

	// Prompt and read the last name
	fmt.Print("Enter last name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName) // Remove any trailing newline

	// Combine first and last name
	fullName := firstName + " " + lastName

	// Output the full name
	fmt.Printf("Name is %s\n", fullName)
}

