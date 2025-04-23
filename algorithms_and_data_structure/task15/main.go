package main

import (
	"fmt"
	"strings"
	
)
//1. Brute Force
// O(n^2) time | O(n) space
func isValid(s string) bool {
	for strings.Contains(s, "()") || strings.Contains(s, "{}") || strings.Contains(s, "[]") {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "[]", "")
	}
	return s == ""
}

func main() {
	// Test cases
	fmt.Println(isValid("[]"))       // true
	fmt.Println(isValid("([{}])"))   // true
	fmt.Println(isValid("[(])"))     // false
	fmt.Println(isValid("{[()]}"))   // true
	fmt.Println(isValid("{[(])}"))   // false
	fmt.Println(isValid("((()))"))   // true
	fmt.Println(isValid("[{}({})]")) // true
}
