package main

import "fmt"

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

func main() {
	// Test cases
	fmt.Println(hammingWeight(0b00000000000000000000000000010111)) // Output: 4
	fmt.Println(hammingWeight(0b01111111111111111111111111111101)) // Output: 30
	fmt.Println(hammingWeight(0b00000000000000000000000010000000)) // Output: 1
	fmt.Println(hammingWeight(0b11111111111111111111111111111111)) // Output: 32
	fmt.Println(hammingWeight(0)) // Output: 0
}
