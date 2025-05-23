//1. Brute Force

package main

import (
	"fmt"
)

func reverseBits(n uint32) uint32 {
	// Convert the number to a binary string.
	binary := ""
	for i := 0; i < 32; i++ {
		if n&(1<<i) != 0 {
			binary += "1"
		} else {
			binary += "0"
		}
	}

	// Reverse the bits using the binary string.
	var res uint32 = 0
	for i, bit := range binary {
		if bit == '1' {
			res |= (1 << (31 - i)) // Shift '1' to the correct position.
		}
	}
	return res
}

func main() {
	// Example 1: Input: n = 21 (binary: 00000000000000000000000000010101)
	n := uint32(21)
	reversed := reverseBits(n)
	fmt.Printf("Original number: %d\n", n)
	fmt.Printf("Reversed number: %d\n", reversed) // Expected: 2818572288 (binary: 10101000000000000000000000000000)

	// Example 2: Input: n = 43261596 (binary: 00000010100101000001111010011100)
	n2 := uint32(43261596)
	reversed2 := reverseBits(n2)
	fmt.Printf("Original number: %d\n", n2)
	fmt.Printf("Reversed number: %d\n", reversed2) // Expected: 964176192 (binary: 00111001011110000010100101000000)
}


//2. Bit Manipulation

func reverseBits1(n uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		bit := (n >> i) & 1
		res |= (bit << (31 - i))
	}
	return res
}

//3. Bit Manipulation (Optimal)

func reverseBits2(n uint32) uint32 {
	res := n
	res = (res >> 16) | (res << 16)
	res = ((res & 0xff00ff00) >> 8) | ((res & 0x00ff00ff) << 8)
	res = ((res & 0xf0f0f0f0) >> 4) | ((res & 0x0f0f0f0f) << 4)
	res = ((res & 0xcccccccc) >> 2) | ((res & 0x33333333) << 2)
	res = ((res & 0xaaaaaaaa) >> 1) | ((res & 0x55555555) << 1)
	return res
}

