package main

import (
	"fmt"
)

// O(n) time | O(n) space
func IsValidParentheses(s string) bool {
	stack := []rune{}
	bracketMap := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range s {
		// If it's a closing bracket
		if openBracket, found := bracketMap[char]; found {
			// Check if the stack is empty or the top of the stack is not the matching opening bracket
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			// Pop from stack
			stack = stack[:len(stack)-1]
		} else {
			// If it's an opening bracket, push it to stack
			stack = append(stack, char)
		}
	}

	// If stack is empty, the string is valid
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


