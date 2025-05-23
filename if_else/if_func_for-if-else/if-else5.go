package main

import (
	"fmt"
	"unicode"
)

// Function to check if a text is a phone number
func isPhoneNumber(text string) bool {
	if len(text) != 12 {
		return false // not phone number-sized
	}

	// Check if the first three characters are digits
	for i := 0; i < 3; i++ {
		if !unicode.IsDigit(rune(text[i])) {
			return false // not an area code
		}
	}

	// Check if the fourth character is a hyphen
	if text[3] != '-' {
		return false // does not have first hyphen
	}

	// Check if the next three characters are digits
	for i := 4; i < 7; i++ {
		if !unicode.IsDigit(rune(text[i])) {
			return false // does not have first 3 digits
		}
	}

	// Check if the eighth character is a hyphen
	if text[7] != '-' {
		return false // does not have second hyphen
	}

	// Check if the last four characters are digits
	for i := 8; i < 12; i++ {
		if !unicode.IsDigit(rune(text[i])) {
			return false // does not have last 4 digits
		}
	}

	return true // text is a phone number
}

func main() {
	// Test cases
	fmt.Println("415-555-4242 is a phone number:")
	fmt.Println(isPhoneNumber("415-555-4242"))
	fmt.Println("Moshi moshi is a phone number:")
	fmt.Println(isPhoneNumber("Moshi moshi"))
}

