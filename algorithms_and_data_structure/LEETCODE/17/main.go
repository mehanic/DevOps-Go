package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	// Edge case: if needle is empty, return 0
	if len(needle) == 0 {
		return 0
	}

	// Loop through haystack up to the point where needle can fit
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// Check if the substring matches needle
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	// If not found
	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))     // Output: 0
	fmt.Println(strStr("leetcode", "leeto"))    // Output: -1
}
