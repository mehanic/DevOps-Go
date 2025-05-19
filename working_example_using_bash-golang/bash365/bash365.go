package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
//	"strings"
)

func main() {
	fmt.Println("### let ###")

	// Bash addition
	addition := 3 + 5
	fmt.Printf("3 + 5 = %d\n", addition)

	// Bash subtraction
	subtraction := 7 - 8
	fmt.Printf("7 - 8 = %d\n", subtraction)

	// Bash multiplication
	multiplication := 5 * 8
	fmt.Printf("5 * 8 = %d\n", multiplication)

	// Bash division
	division := 4 / 2
	fmt.Printf("4 / 2 = %d\n", division)

	// Bash modulus
	modulus := 9 % 4
	fmt.Printf("9 %% 4 = %d\n", modulus)

	// Bash power of two
	powerOfTwo := int(math.Pow(2, 2))
	fmt.Printf("2 ^ 2 = %d\n", powerOfTwo)

	fmt.Println("### Bash Arithmetic Expansion ###")

	// Arithmetic expressions
	fmt.Printf("4 + 5 = %d\n", 4+5)
	fmt.Printf("7 - 7 = %d\n", 7-7)
	fmt.Printf("4 x 6 = %d\n", 3*2)
	fmt.Printf("6 / 3 = %d\n", 6/3)
	fmt.Printf("8 %% 7 = %d\n", 8%7)
	fmt.Printf("2 ^ 8 = %d\n", int(math.Pow(2, 8)))

	fmt.Println("### Declare ###")

	// User input
	var num1, num2 int
	fmt.Print("Please enter two numbers: ")
	_, err := fmt.Scanf("%d %d", &num1, &num2)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// Calculation of result
	result := num1 + num2
	fmt.Printf("Result is: %d\n", result)

	// Binary conversion (10001)
	binaryStr := "10001"
	binaryNum, _ := strconv.ParseInt(binaryStr, 2, 64)
	fmt.Printf("Binary %s as decimal is: %d\n", binaryStr, binaryNum)

	// Octal conversion (16)
	octalStr := "16"
	octalNum, _ := strconv.ParseInt(octalStr, 8, 64)
	fmt.Printf("Octal %s as decimal is: %d\n", octalStr, octalNum)

	// Hexadecimal conversion (0xE6A)
	hexStr := "E6A"
	hexNum, _ := strconv.ParseInt(hexStr, 16, 64)
	fmt.Printf("Hexadecimal %s as decimal is: %d\n", hexStr, hexNum)
}

