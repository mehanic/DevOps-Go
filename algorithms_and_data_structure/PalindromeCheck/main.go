package main

import (
	"fmt"
)

// O(n^2) time | O(n) space
func IsPalindrome(str string) bool {
	reversed := make([]byte, len(str)) // Create a new slice to store the reversed string

	// Reverse the string
	for i := len(str) - 1; i >= 0; i-- {
		j := len(str) - i - 1
		reversed[j] = str[i]
	}

	// Compare original and reversed strings
	for i := range str {
		if reversed[i] != str[i] {
			return false
		}
	}
	return true
}

func main() {
	// Example test cases
	words := []string{"madam", "racecar", "hello", "level", "world"}

	for _, word := range words {
		fmt.Printf("Is \"%s\" a palindrome? %v\n", word, IsPalindrome(word))
	}

	fmt.Println("-----")
	main1()
	fmt.Println("-----")
	main2()
	fmt.Println("-----")
	main3()

}

///////

// O(n) time | O(n) space
func IsPalindrome1(str string) bool {
	result := []byte{} // Create an empty slice to store the reversed string

	// Reverse the string
	for i := len(str) - 1; i >= 0; i-- {
		result = append(result, str[i])
	}

	// Convert the reversed slice to a string and compare it with the original
	return str == string(result)
}

func main1() {
	// Test cases
	words := []string{"madam", "racecar", "hello", "level", "world"}

	for _, word := range words {
		fmt.Printf("Is \"%s\" a palindrome? %v\n", word, IsPalindrome1(word))
	}
}

///////

// O(n) time | O(n) space
func IsPalindrome2(str string) bool {
	return helper(str, 0)
}

func helper(str string, i int) bool {
	j := len(str) - 1 - i
	if i >= j {
		return true
	}
	if str[i] != str[j] {
		return false
	}
	return helper(str, i+1)
}

func main2() {
	// Test cases
	words := []string{"madam", "racecar", "hello", "level", "world"}

	for _, word := range words {
		fmt.Printf("Is \"%s\" a palindrome? %v\n", word, IsPalindrome2(word))
	}
}

////

// O(n) time | O(1) space
func IsPalindrome3(str string) bool {
	for i := 0; i < len(str)/2; i++ { // Iterate only till the middle
		j := len(str) - i - 1 // Calculate corresponding right index
		if str[i] != str[j] {
			return false // If characters don't match, return false
		}
	}
	return true // If loop completes, it's a palindrome
}

func main3() {
	// Test cases
	words := []string{"madam", "racecar", "hello", "level", "world"}

	for _, word := range words {
		fmt.Printf("Is \"%s\" a palindrome? %v\n", word, IsPalindrome(word))
	}
}
