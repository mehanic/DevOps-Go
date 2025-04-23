package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
    n := len(digits)
    
    for i := n - 1; i >= 0; i-- {
        if digits[i] < 9 {
            digits[i]++ // Increment the digit
            return digits
        }
        digits[i] = 0 // Carry over, set to 0
    }

    // If we exit the loop, that means all digits were 9
    return append([]int{1}, digits...) // Insert 1 at the beginning
}

func main() {
    fmt.Println(plusOne([]int{1, 2, 3, 4})) // Output: [1, 2, 3, 5]
    fmt.Println(plusOne([]int{9, 9, 9}))    // Output: [1, 0, 0, 0]
    fmt.Println(plusOne([]int{4, 5, 9}))    // Output: [4, 6, 0]
}
