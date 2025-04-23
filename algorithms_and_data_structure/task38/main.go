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
    fmt.Println(isHappy(19)) // Output: true
    fmt.Println(isHappy(2))  // Output: false
    fmt.Println(isHappy(7))  // Output: true
}


//---
//2. Fast And Slow Pointers - I

func isHappy1(n int) bool {
    slow, fast := n, sumOfSquares(n)

    for slow != fast {
        fast = sumOfSquares(fast)
        fast = sumOfSquares(fast)
        slow = sumOfSquares(slow)
    }

    return fast == 1
}

func sumOfSquares1(n int) int {
    output := 0
    for n > 0 {
        digit := n % 10
        output += digit * digit
        n = n / 10
    }
    return output
}

//3. Fast And Slow Pointers - II
func isHappy2(n int) bool {
    slow, fast := n, sumOfSquares(n)
    power, lam := 1, 1

    for slow != fast {
        if power == lam {
            slow = fast
            power *= 2
            lam = 0
        }
        fast = sumOfSquares(fast)
        lam++
    }

    return fast == 1
}

func sumOfSquares2(n int) int {
    output := 0
    for n > 0 {
        digit := n % 10
        output += digit * digit
        n = n / 10
    }
    return output
}

