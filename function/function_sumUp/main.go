package main

import (
	"fmt"
)

// sumUpTo calculates the sum of all numbers from 1 up to (but not including) n
func sumUpTo(n int) int {
	return n * (n - 1) / 2
}

func main() {
	// Example usage
	n := 10
	result := sumUpTo(n)
	fmt.Printf("The sum of numbers from 1 to %d (excluding %d) is: %d\n", n-1, n, result)
}

