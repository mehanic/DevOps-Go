package main

import (
	"fmt"
	"unicode"
)

// checkio returns the most common letter in the input text.
func checkio(text string) string {
	// Create a map to count occurrences of each letter
	counts := make(map[rune]int)

	// Normalize the text: lower-case and count letters
	for _, char := range text {
		// Convert to lowercase and check if it's a letter
		lowerChar := unicode.ToLower(char)
		if unicode.IsLetter(lowerChar) {
			counts[lowerChar]++
		}
	}

	// Variables to track the most common letter and its count
	var mostCommonChar rune
	maxCount := 0

	// Find the most common letter
	for char, count := range counts {
		// Check for higher count or tie-breaking on the letter order
		if count > maxCount || (count == maxCount && char < mostCommonChar) {
			mostCommonChar = char
			maxCount = count
		}
	}

	return string(mostCommonChar)
}

func main() {
	// Test cases to validate the function
	if checkio("Hello World!") != "l" {
		fmt.Println("Hello test failed")
	}
	if checkio("How do you do?") != "o" {
		fmt.Println("O is most wanted test failed")
	}
	if checkio("One") != "e" {
		fmt.Println("All letter only once test failed")
	}
	if checkio("Oops!") != "o" {
		fmt.Println("Don't forget about lower case test failed")
	}
	if checkio("AAaooo!!!!") != "a" {
		fmt.Println("Only letters test failed")
	}

	fmt.Println("All tests passed!")
}

