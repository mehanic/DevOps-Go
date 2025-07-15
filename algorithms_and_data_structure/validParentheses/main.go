package main

import (
	"fmt"
)

func IsValidParentheses(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		if openBracket, found := bracketMap[char]; found {
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}

func main() {
	// Test cases
	fmt.Println(IsValidParentheses("[]"))       // Output: true
	fmt.Println(IsValidParentheses("([{}])"))   // Output: true
	fmt.Println(IsValidParentheses("[(])"))     // Output: false
	fmt.Println(IsValidParentheses("{[()]}"))   // Output: true
	fmt.Println(IsValidParentheses("{[(])}"))   // Output: false
	fmt.Println(IsValidParentheses("((()))"))   // Output: true
	fmt.Println(IsValidParentheses("[{}({})]")) // Output: true
}


