package main

import "fmt"

// A loop that prints iteration numbers and stops at 5
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Iteration number", i)
		if i == 5 {
			break // Exit the loop when `i` equals 5
		}
	}

	main1()
}

// A loop that skips even numbers and prints odd numbers
func main1() {
	for i := 0; i < 10; i++ {
		// If the result is even, skip to the next iteration
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "is an odd number")
	}
}


// Iteration number 0
// Iteration number 1
// Iteration number 2
// Iteration number 3
// Iteration number 4
// Iteration number 5
// 1 is an odd number
// 3 is an odd number
// 5 is an odd number
// 7 is an odd number
// 9 is an odd number
