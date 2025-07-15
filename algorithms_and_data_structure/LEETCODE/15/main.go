package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	var filtered []rune
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	n := len(filtered)
	for i := 0; i < n/2; i++ {
		if filtered[i] != filtered[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false

	// TODO: implement
}
