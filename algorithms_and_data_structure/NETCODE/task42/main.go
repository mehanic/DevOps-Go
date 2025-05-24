package main

import (
	"fmt"
	"math/bits"
)

// Function to count the number of 1 bits in the binary representation of a number
func hammingWeight(n int) int {
	res := 0
	// Iterate through all 32 bits
	for i := 0; i < 32; i++ {
		// Check if the i-th bit is set to 1
		if (1<<i)&n != 0 {
			res++
		}
	}
	return res
}
//1. Bit Mask - I
func main() {
	// Test cases
	fmt.Println(hammingWeight(0b00000000000000000000000000010111)) // Output: 4
	fmt.Println(hammingWeight(0b01111111111111111111111111111101)) // Output: 30
	fmt.Println(hammingWeight(0b00000000000000000000000010000000)) // Output: 1
	fmt.Println(hammingWeight(0b11111111111111111111111111111111)) // Output: 32
	fmt.Println(hammingWeight(0)) // Output: 0
}


//2. Bit Mask - II

func hammingWeight1(n int) int {
	res := 0
	for n != 0 {
		if n&1 != 0 {
			res++
		}
		n >>= 1
	}
	return res
}


//3. Bit Mask (Optimal)

func hammingWeight2(n int) int {
	res := 0
	for n != 0 {
		n &= n - 1
        res++
	}
	return res
}

//4. Built-In Function

func hammingWeight3(n int) int {
	return bits.OnesCount(uint(n))
}


//

