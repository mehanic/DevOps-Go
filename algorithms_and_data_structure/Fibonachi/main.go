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

}
