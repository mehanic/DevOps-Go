package main

import (
    "fmt"
)

func isHappy(n int) bool {
    visit := make(map[int]bool)

    for !visit[n] {
        visit[n] = true
        n = sumOfSquares(n)
        if n == 1 {
            return true
        }
    }
    return false
}

func sumOfSquares(n int) int {
    output := 0
    for n > 0 {
        digit := n % 10
        digit = digit * digit
        output += digit
        n = n / 10
    }
    return output
}

func main() {
    fmt.Println("Is 19 a happy number?", isHappy(19))  // true
    fmt.Println("Is 2 a happy number?", isHappy(2))    // false
}
