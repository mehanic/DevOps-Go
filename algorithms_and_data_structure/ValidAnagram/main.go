// Valid Anagram
// Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

// An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

// Example 1:

// Input: s = "racecar", t = "carrace"

// Output: true
// Example 2:

// Input: s = "jar", t = "jam"

// Output: false
// Constraints:

// s and t consist of lowercase English letters.


// Recommended Time & Space Complexity



package main

import (
	"fmt"
)
//1. Sorting
func isAnagram(s string, t string) bool {
	// If lengths are different, they cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Create a frequency map to count character occurrences
	count := make(map[rune]int)

	// Count characters in the first string
	for _, char := range s {
		count[char]++
	}

	// Subtract character counts using the second string
	for _, char := range t {
		count[char]--
		// If a count becomes negative, t has an extra character
		if count[char] < 0 {
			return false
		}
	}

	return true
}

func main() {
	// Test cases
	fmt.Println(isAnagram("racecar", "carrace")) // true
	fmt.Println(isAnagram("jar", "jam"))         // false
	fmt.Println(isAnagram("listen", "silent"))   // true
	fmt.Println(isAnagram("hello", "world"))     // false
}
