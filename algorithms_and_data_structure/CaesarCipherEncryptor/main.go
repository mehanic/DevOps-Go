package main

import (
	"fmt"
	"strings"
)

// O(n) time | O(n) space
func CaesarCipherEncryptor(str string, key int) string {
	shift := rune(key % 26) // Ensure key stays within 0-25
	runes := []rune(str)

	for i, char := range runes {
		if char >= 'a' && char <= 'z' { // Ensure it's a lowercase letter
			newChar := char + shift
			if newChar > 'z' {
				newChar -= 26 // Wrap around the alphabet
			}
			runes[i] = newChar
		}
	}
	return string(runes)
}

func main() {
	// Test cases
	fmt.Println(CaesarCipherEncryptor("abc", 2))   // Output: "cde"
	fmt.Println(CaesarCipherEncryptor("xyz", 5))   // Output: "cde"
	fmt.Println(CaesarCipherEncryptor("hello", 3)) // Output: "khoor"
	fmt.Println(CaesarCipherEncryptor("zebra", 1)) // Output: "afcsb"
	main1()
}


//////---


// O(n) time | O(n) space
func CaesarCipherEncryptor1(str string, key int) string {
	runes := []rune(str)
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i, char := range runes {
		index := strings.Index(alphabet, string(char))
		if index == -1 {
			return "" // Bad input (non-lowercase letters)
		}
		newIndex := (index + key) % 26
		runes[i] = rune(alphabet[newIndex])
	}
	return string(runes)
}

func main1() {
	// Test cases
	fmt.Println(CaesarCipherEncryptor("abc", 2))   // Output: "cde"
	fmt.Println(CaesarCipherEncryptor("xyz", 5))   // Output: "cde"
	fmt.Println(CaesarCipherEncryptor("hello", 3)) // Output: "khoor"
	fmt.Println(CaesarCipherEncryptor("zebra", 1)) // Output: "afcsb"
}
