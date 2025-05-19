package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("---- let ----")
	addition := 1 + 2
	fmt.Printf("1 + 2 = %d\n", addition)

	subtraction := 4 - 3
	fmt.Printf("4 - 3 = %d\n", subtraction)

	multiplication := 2 * 3
	fmt.Printf("2 * 3 = %d\n", multiplication)

	division := 4 / 2
	fmt.Printf("4 / 2 = %d\n", division)

	mod := 4 % 3
	fmt.Printf("4 %% 3 = %d\n", mod)

	powerOfTwo := int(math.Pow(2, 2)) // 2 to the power of 2
	fmt.Printf("2 ** 2 = %d\n", powerOfTwo)

	fmt.Println("---- Go Arithmetic Expansion ----")
	fmt.Printf("1 + 2 = %d\n", 1+2)
	fmt.Printf("4 - 3 = %d\n", 4-3)
	fmt.Printf("2 * 3 = %d\n", 2*3)
	fmt.Printf("4 / 2 = %d\n", 4/2)
	fmt.Printf("4 %% 3 = %d\n", 4%3)
	fmt.Printf("2 ** 2 = %d\n", int(math.Pow(2, 2)))

	var n1, n2 int
	fmt.Print("Please enter two numbers: ")
	fmt.Scan(&n1, &n2)

	result := n1 + n2
	fmt.Printf("result is: %d\n", result)

	// Binary, Octal, and Hexadecimal
	binaryResult := 2 << 3 // Equivalent to binary 1001 (2^3 = 8)
	fmt.Printf("binary: 1001 -> decimal: %d\n", binaryResult)

	octalResult := 011 // Go interprets this as octal
	fmt.Printf("octal: 11 -> decimal: %d\n", octalResult)

	hexResult := 0xF // Hexadecimal F
	fmt.Printf("hex: F -> decimal: %d\n", hexResult)
}

