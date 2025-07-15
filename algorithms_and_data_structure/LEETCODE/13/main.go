package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	number := 0
	for x != 0 {
		digit := x % 10
		x /= 10

		// Overflow check
		if number > (math.MaxInt32-digit)/10 {
			// return 0
		}

		number = number*10 + digit
	}

	return sign * number
}

func reverse1(x int) int {
	const INT_MAX = 1<<31 - 1 // 2147483647
	const INT_MIN = -1 << 31  // -2147483648

	result := 0

	for x != 0 {
		digit := x % 10
		x /= 10

		// Check for overflow before it happens
		if result > (INT_MAX-digit)/10 || result < (INT_MIN-digit)/10 {
			return 0
		}

		result = result*10 + digit
	}

	return result
}

func main() {
	num1 := 24
	result1 := reverse(num1)
	fmt.Printf("reverse(%d) = %d\n", num1, result1)

	num2 := 56
	result2 := reverse1(num2)
	fmt.Printf("reverse(%d) = %d\n", num2, result2) // TODO: implement
}
