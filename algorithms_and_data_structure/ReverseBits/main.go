// Reverse Bits
// Given a 32-bit unsigned integer n, reverse the bits of the binary representation of n and return the result.

// Example 1:

// Input: n = 00000000000000000000000000010101

// Output:    2818572288 (10101000000000000000000000000000)
// Explanation: Reversing 00000000000000000000000000010101, which represents the unsigned integer 21, gives us 10101000000000000000000000000000 which represents the unsigned integer 2818572288.




package main

import "fmt"

// reverseBits reverses the bits of a 32-bit unsigned integer.
func reverseBits(n uint32) uint32 {
    var result uint32 = 0
    for i := 0; i < 32; i++ {
        result <<= 1           // Shift result left by 1 to make space for the next bit.
        result |= n & 1        // Add the least significant bit of n to result.
        n >>= 1                // Right shift n to process the next bit.
    }
    return result
}

func main() {
    // Example 1
    n := uint32(21) // Binary: 00000000000000000000000000010101
    reversed := reverseBits(n)
    fmt.Println("Reversed Bits:", reversed) // Expected: 2818572288 (Binary: 10101000000000000000000000000000)
}
