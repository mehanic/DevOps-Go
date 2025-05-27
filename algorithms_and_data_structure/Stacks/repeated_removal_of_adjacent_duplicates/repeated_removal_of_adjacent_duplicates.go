package main

import (
	"fmt"
)

func repeatedRemovalOfAdjacentDuplicates(s string) string {
	stack := []rune{}

	for _, c := range s {
		n := len(stack)
		if n > 0 && c == stack[n-1] {
			stack = stack[:n-1] // удаляем последний элемент (пару дубликатов)
		} else {
			stack = append(stack, c)
		}
	}

	return string(stack)
}

func main() {
	s := "abbaca"
	fmt.Println(repeatedRemovalOfAdjacentDuplicates(s)) // Output: "ca"
}
