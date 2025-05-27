package main

import (
	"fmt"
)

func longestSubstringWithUniqueCharsOptimized(s string) int {
	maxLen := 0
	prevIndexes := make(map[rune]int)
	left := 0
	runes := []rune(s) // Используем руны для поддержки Unicode

	for right, char := range runes {
		if prevIdx, exists := prevIndexes[char]; exists && prevIdx >= left {
			left = prevIdx + 1
		}
		if currentLen := right - left + 1; currentLen > maxLen {
			maxLen = currentLen
		}
		prevIndexes[char] = right
	}

	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Println("Longest substring with unique characters:", longestSubstringWithUniqueCharsOptimized(s)) // Output: 3
}
