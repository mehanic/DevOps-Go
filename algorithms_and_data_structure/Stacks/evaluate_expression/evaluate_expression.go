package main

import (
	"fmt"
	"unicode"
)

func evaluateExpression(s string) int {
	stack := []int{}
	currNum := 0
	sign := 1
	res := 0

	for _, c := range s {
		if unicode.IsDigit(c) {
			currNum = currNum*10 + int(c-'0')
		} else if c == '+' || c == '-' {
			res += sign * currNum
			if c == '-' {
				sign = -1
			} else {
				sign = 1
			}
			currNum = 0
		} else if c == '(' {
			// Push current result and sign onto stack
			stack = append(stack, res)
			stack = append(stack, sign)
			// Reset res and sign for new expression inside parentheses
			res = 0
			sign = 1
		} else if c == ')' {
			// Complete current nested expression
			res += sign * currNum
			currNum = 0
			// Apply sign from stack
			if len(stack) > 0 {
				signFromStack := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res *= signFromStack
			}
			// Add result from stack
			if len(stack) > 0 {
				prevRes := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				res += prevRes
			}
		}
	}

	return res + sign*currNum
}

func main() {
	expr := "(1+(4+5+2)-3)+(6+8)"
	fmt.Println(evaluateExpression(expr)) // Output: 23
}
