package main

import (
	"fmt"
	"strings"
)

func main() {
	// Define the input string
	input := "Good Morning and have a nice day !"

	// Split the string into words using spaces as the delimiter
	words := strings.Fields(input)

	// Count the number of words
	wordCount := len(words)

	// Print the number of words
	fmt.Printf("The number of words is: %d\n", wordCount)
}

