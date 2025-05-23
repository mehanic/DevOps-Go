package main

import (
	"fmt"
//	"math/big"
	"math"
)

// Function to perform precision scale operations
func precisionScale() {
	result := 3.0 / 8.0
	fmt.Printf("Precision Scale (3/8): %.2f\n", result)
}

// Function to convert decimal to binary
func decimalToBinary(no int) {
	binary := fmt.Sprintf("%b", no)
	fmt.Printf("Decimal %d to Binary: %s\n", no, binary)
}

// Function to convert binary to decimal
func binaryToDecimal(binaryStr string) {
	var decimal int64
	for i := 0; i < len(binaryStr); i++ {
		decimal = decimal*2 + int64(binaryStr[i]-'0')
	}
	fmt.Printf("Binary %s to Decimal: %d\n", binaryStr, decimal)
}

// Function to calculate square root
func squareRoot(value float64) {
	result := math.Sqrt(value)
	fmt.Printf("Square Root of %.2f: %.2f\n", value, result)
}

// Function to calculate power
func power(base, exponent float64) {
	result := math.Pow(base, exponent)
	fmt.Printf("%.2f to the power of %.2f: %.0f\n", base, exponent, result)
}

func main() {
	no := 100

	// Set precision scale
	precisionScale()

	// Convert decimal to binary
	decimalToBinary(no)

	// Convert binary to decimal (from binary string)
	binaryNo := "1100100"
	binaryToDecimal(binaryNo)

	// Calculate square root
	squareRoot(100)

	// Calculate power
	power(10, 10)
}

