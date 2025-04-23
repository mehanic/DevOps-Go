package main

import (
	"fmt"
	
)
//1. Reverse String
func isPalindrome(s string) bool {
	newStr := ""
	for _, c := range s {
		if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
			newStr += string(c) // Keep lowercase letters and numbers
		} else if 'A' <= c && c <= 'Z' {
			newStr += string(c + 'a' - 'A') // Convert uppercase to lowercase
		}
	}

	reversedStr := reverse(newStr) // Reverse the cleaned string
	return newStr == reversedStr
}

func reverse(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i] // Swap elements
	}
	return string(runes)
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
	fmt.Println("-----------")
}