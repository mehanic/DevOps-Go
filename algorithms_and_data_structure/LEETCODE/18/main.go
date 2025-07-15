package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0] // Start with the first string as a prefix

	for _, s := range strs[1:] {
		// Shrink the prefix until it matches the start of s
		for len(prefix) > 0 && len(s) < len(prefix) || s[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))  // Output: "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // Output: ""
}
