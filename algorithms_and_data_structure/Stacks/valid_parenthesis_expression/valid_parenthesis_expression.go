package main

import "fmt"

func validParenthesisExpression(s string) bool {
    parenthesesMap := map[rune]rune{
        '(': ')',
        '{': '}',
        '[': ']',
    }
    stack := []rune{}

    for _, c := range s {
        if _, ok := parenthesesMap[c]; ok {
            // Если открывающая скобка — кладём в стек
            stack = append(stack, c)
        } else {
            // Иначе проверяем закрывающую
            n := len(stack)
            if n == 0 {
                return false
            }
            last := stack[n-1]
            if parenthesesMap[last] == c {
                stack = stack[:n-1]
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}

func main() {
    fmt.Println(validParenthesisExpression("()[]{}"))   // true
    fmt.Println(validParenthesisExpression("([)]"))     // false
    fmt.Println(validParenthesisExpression("{[]}"))     // true
}
