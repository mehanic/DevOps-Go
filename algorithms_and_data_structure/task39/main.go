package main

import (
	"fmt"
)

// 1. Recursion
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	if digits[len(digits)-1] < 9 {
		digits[len(digits)-1]++
		return digits
	} else {
		return append(plusOne(digits[:len(digits)-1]), 0)
	}
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3, 4})) // Output: [1, 2, 3, 5]
	fmt.Println(plusOne([]int{9, 9, 9}))    // Output: [1, 0, 0, 0]
	fmt.Println(plusOne([]int{4, 5, 9}))    // Output: [4, 6, 0]
	fmt.Println(plusOne([]int{0}))          // Output: [1]
}

// 2. Iteration - I
func plusOne1(digits []int) []int {
	one := 1
	i := 0
	digits = reverse(digits)

	for one != 0 {
		if i < len(digits) {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i] += 1
				one = 0
			}
		} else {
			digits = append(digits, one)
			one = 0
		}
		i++
	}
	return reverse(digits)
}

func reverse(digits []int) []int {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}

/////

// 3. Iteration - II
func plusOne2(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

//
