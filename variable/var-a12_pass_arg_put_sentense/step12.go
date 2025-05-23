package main

import "fmt"

func main() {

	var number int
	fmt.Print("Enter a number: ") // Prompt user for input
	fmt.Scanln(&number)           // Read input

	// Output the result of multiplying the input by 5
	fmt.Printf("When you multiply %d by 5, it becomes %d\n", number, number*5)
}

// Example:
// Enter a number: 56
// When you multiply 56 by 5, it becomes 280
