package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a 'list'.
	words := "house dog telephone dog"

	// Split the words into a slice.
	wordList := strings.Fields(words)

	// Iterate over the list and process the values.
	for _, word := range wordList {
		fmt.Printf("The word is: %s\n", word)
	}
}

