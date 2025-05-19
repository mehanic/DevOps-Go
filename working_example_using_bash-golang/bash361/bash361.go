package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	// Prompt the user for input
	fmt.Print("Please Enter the String: ")
	fmt.Scanln(&str) // Read input from user

	// Calculate the length of the string
	len := len(strings.TrimSpace(str)) // Using TrimSpace to remove extra spaces

	// Print the length of the string
	fmt.Printf("Length of string = %d\n", len)
}

