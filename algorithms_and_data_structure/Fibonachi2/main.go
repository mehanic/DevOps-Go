package main

import "fmt"

func GetNthFib(n int) int {
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

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, GetNthFib(i))
	}
}
