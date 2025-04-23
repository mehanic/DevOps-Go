package main

import "fmt"

// O(2^n) time | O(n) space
func GetNthFib(n int) int {
	if n == 2 {
		return 1
	} else if n == 1 {
		return 0
	}
	return GetNthFib(n-1) + GetNthFib(n-2)
}

func main() {
	// Example: Print the first 10 Fibonacci numbers
	for i := 1; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, GetNthFib(i))
	}
	fmt.Println("-------")
	main1()
	fmt.Println("-------")
	main2()
}

// ////
// O(n) time | O(n) space (due to recursion and memoization)
func GetNthFib1(n int) int {
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

	// If the value is already computed, return it
	if val, found := memoize[n]; found {
		return val
	}

	// Compute and store the result
	memoize[n] = helper(n-1, memoize) + helper(n-2, memoize)
	return memoize[n]
}

func main1() {
	// Example: Print the first 10 Fibonacci numbers
	for i := 1; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, GetNthFib1(i))
	}
}

// O(n) time | O(1) space
func GetNthFib2(n int) int {
	lastTwo := []int{0, 1} // Stores the last two Fibonacci numbers
	counter := 3           // Start from the 3rd Fibonacci number

	// Iteratively compute Fibonacci numbers up to n
	for counter <= n {
		nextFib := lastTwo[0] + lastTwo[1]   // Compute next Fibonacci number
		lastTwo = []int{lastTwo[1], nextFib} // Update last two numbers
		counter += 1
	}

	// Return the correct Fibonacci number
	if n > 1 {
		return lastTwo[1] // If n > 1, return the latest computed Fibonacci number
	}
	return lastTwo[0] // If n == 1, return 0
}

func main2() {
	// Test the function with different Fibonacci numbers
	for i := 1; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, GetNthFib(i))
	}
}
