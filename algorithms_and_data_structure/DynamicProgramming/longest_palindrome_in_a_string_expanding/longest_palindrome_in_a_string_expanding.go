package main

import (
	"fmt"
)

func longestPalindromeInAStringExpanding(s string) string {
	n := len(s)
	start, maxLen := 0, 0

	for center := 0; center < n; center++ {
		// Odd-length palindrome
		oddStart, oddLen := expandPalindrome(center, center, s)
		if oddLen > maxLen {
			start = oddStart
			maxLen = oddLen
		}
		// Even-length palindrome
		if center < n-1 && s[center] == s[center+1] {
			evenStart, evenLen := expandPalindrome(center, center+1, s)
			if evenLen > maxLen {
				start = evenStart
				maxLen = evenLen
			}
		}
	}
	return s[start : start+maxLen]
}

// Expand from the center and return the start and length of the palindrome.
func expandPalindrome(left, right int, s string) (int, int) {
	for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
		left--
		right++
	}
	return left, right - left + 1
}

func main() {
	s := "babad"
	fmt.Println(longestPalindromeInAStringExpanding(s)) // Output: "bab" or "aba"
}
