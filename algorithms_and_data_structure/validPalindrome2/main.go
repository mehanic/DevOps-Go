package main

import (
	"fmt"
	"unicode"
)

// Two Pointers
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func main() {
	// Test cases
	examples := []string{
		"Was it a car or a cat I saw?",    // True
		"A man, a plan, a canal: Panama!", // True
		"racecar",                         // True
		"hello",                           // False
		"tab a cat",                       // False
	}

	for _, example := range examples {
		fmt.Printf("isPalindrome(%q) = %v\n", example, isPalindrome(example))
	}
}
