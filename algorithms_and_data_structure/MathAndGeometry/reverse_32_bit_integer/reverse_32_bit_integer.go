package main

import (
	"fmt"
	"math"
)

func reverse32BitInteger(n int) int {
	const INT_MAX = math.MaxInt32
	const INT_MIN = math.MinInt32
	reversed := 0

	for n != 0 {
		digit := n % 10
		n = n / 10

		// Check for overflow/underflow before multiplying by 10
		if reversed > INT_MAX/10 || (reversed == INT_MAX/10 && digit > 7) {
			return 0
		}
		if reversed < INT_MIN/10 || (reversed == INT_MIN/10 && digit < -8) {
			return 0
		}

		reversed = reversed*10 + digit
	}

	return reversed
}

func main() {
	fmt.Println(reverse32BitInteger(123))       // 321
	fmt.Println(reverse32BitInteger(-123))      // -321
	fmt.Println(reverse32BitInteger(1534236469)) // 0 (overflow)
}
