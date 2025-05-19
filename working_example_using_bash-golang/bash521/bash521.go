package main

import (
	"fmt"
)

func add(a, b int) (int, int) {
	sum := a + b
	dif := a - b
	return sum, dif
}

func main() {
	var int1, int2 int

	// Prompt the user for input
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&int1)
	fmt.Print("Enter an integer: ")
	fmt.Scanln(&int2)

	// Call the add function
	sum, dif := add(int1, int2)

	// Display the results
	fmt.Println("The sum is:", sum)
	fmt.Println("The difference is:", dif)
}

