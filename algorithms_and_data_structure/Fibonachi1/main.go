package main

import "fmt"

func GetNthFib(n int) int {
	memoize := make(map[int]int) // Create a map to store computed Fibonacci values
	return helper(n, memoize)
}

// Helper function with memoization
func helper(n int, memoize map[int]int) int {
	// Base cases
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	}

	if val, found := memoize[n]; found {
		return val
	}

	memoize[n] = helper(n-1, memoize) + helper(n-2, memoize)
	return memoize[n]
}

func main() {
	// Example: Print the first 10 Fibonacci numbers
	for i := 1; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, GetNthFib(i))
	}
}
