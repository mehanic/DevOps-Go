// Number of One Bits
// You are given an unsigned integer n. Return the number of 1 bits in its binary representation.

// You may assume n is a non-negative integer which fits within 32-bits.

// Example 1:

// Input: n = 00000000000000000000000000010111

// Output: 4
// Example 2:

// Input: n = 01111111111111111111111111111101

// Output: 30



package main

import "fmt"

// Function to count the number of 1 bits
func hammingWeight(n uint32) int {
    count := 0
    for n != 0 {
        n &= (n - 1)  // This operation removes the rightmost '1' bit from n
        count++
    }
    return count
}

func main() {
    // Example 1: Test input 00000000000000000000000000010111
    n1 := uint32(0b00000000000000000000000000010111)
    fmt.Println(hammingWeight(n1)) // Output: 4

    // Example 2: Test input 01111111111111111111111111111101
    n2 := uint32(0b01111111111111111111111111111101)
    fmt.Println(hammingWeight(n2)) // Output: 30
}

