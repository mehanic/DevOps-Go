
package main

import (
	"fmt"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	
)

func isValid1(s string) bool {
    stack := linkedliststack.New()
    closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

    for _, c := range s {
        if open, exists := closeToOpen[c]; exists {
            if !stack.Empty() {
                top, ok := stack.Pop()
                if ok && top.(rune) != open {
                    return false
                }
            } else {
                return false
            }
        } else {
            stack.Push(c)
        }
    }

    return stack.Empty()
}

func main() {
	// Test cases
	examples := []string{"()", "()[]{}", "(]", "([)]", "{[]}"}

	for _, s := range examples {
		fmt.Printf("isValid(%q) = %v\n", s, isValid1(s))
	}
}
