
// Non-Cyclical Number
// A non-cyclical number is an integer defined by the following algorithm:

// Given a positive integer, replace it with the sum of the squares of its digits.
// Repeat the above step until the number equals 1, or it loops infinitely in a cycle which does not include 1.
// If it stops at 1, then the number is a non-cyclical number.
// Given a positive integer n, return true if it is a non-cyclical number, otherwise return false.

// Example 1:

// Input: n = 100

// Output: true
// Explanation: 1^2 + 0^2 + 0^2 = 1

// Example 2:

// Input: n = 101

// Output: false
// Explanation:
// 1^2 + 0^2 + 1^2 = 2
// 2^2 = 4
// 4^2 = 16
// 1^2 + 6^2 = 37
// 3^2 + 7^2 = 58
// 5^2 + 8^2 = 89
// 8^2 + 9^2 = 145
// 1^2 + 4^2 + 5^2 = 42
// 4^2 + 2^2 = 20
// 2^2 + 0^2 = 4 (This number has already been seen)

// Constraints:

// 1 <= n <= 1000


package main

import (
	"fmt"
)

// Function to check if a number is non-cyclical
func isNonCyclicalNumber(n int) bool {
	seen := make(map[int]bool) // Stores numbers seen before

	for n != 1 {
		if seen[n] {
			return false // Loop detected
		}
		seen[n] = true
		n = sumOfSquares(n) // Compute next value
	}
	return true
}

// Helper function to calculate the sum of squares of digits
func sumOfSquares(n int) int {
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += digit * digit
		n /= 10
	}
	return sum
}

func main() {
	// Example 1: n = 100
	n1 := 100
	fmt.Printf("Is %d a non-cyclical number? %v\n", n1, isNonCyclicalNumber(n1)) // Output: true

	// Example 2: n = 101
	n2 := 101
	fmt.Printf("Is %d a non-cyclical number? %v\n", n2, isNonCyclicalNumber(n2)) // Output: false

	// Additional Test Cases
	fmt.Printf("Is %d a non-cyclical number? %v\n", 19, isNonCyclicalNumber(19)) // Output: true
	fmt.Printf("Is %d a non-cyclical number? %v\n", 2, isNonCyclicalNumber(2))   // Output: false
}
