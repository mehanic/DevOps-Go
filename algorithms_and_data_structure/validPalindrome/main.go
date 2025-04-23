package main

import (
	"fmt"
	"unicode"
)

// isPalindrome checks if a given string is a valid palindrome
func isPalindrome(s string) bool {
	var filtered []rune

	// Step 1: Filter the string (remove non-alphanumeric & convert to lowercase)
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	// Step 2: Two-pointer check
	left, right := 0, len(filtered)-1
	for left < right {
		if filtered[left] != filtered[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	// Test cases
	examples := []string{
		"Was it a car or a cat I saw?",    // true
		"tab a cat",                       // false
		"A man, a plan, a canal: Panama!", // true
		"RaceCar",                         // true
		"hello",                           // false
	}

	// Run test cases
	for _, example := range examples {
		fmt.Println("isPalindrome(%q) = %v\n", example, isPalindrome(example))
	}
}
