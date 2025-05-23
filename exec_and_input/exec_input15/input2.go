package main

import (
	"fmt"
	"strconv"
)

// Function to calculate and print the addition
func getAdd(p, q int) {
	aresult := p + q
	fmt.Printf("The addition of %d and %d is: %d\n", p, q, aresult)
}

// Function to calculate and print the subtraction
func getSub(m, n int) {
	sresult := m - n
	fmt.Printf("The sub of %d and %d is: %d\n", m, n, sresult)
}

func main() {
	// Prompt user for input
	fmt.Print("Enter your first num: ")
	var input1 string
	fmt.Scanln(&input1)
	
	fmt.Print("Enter your second num: ")
	var input2 string
	fmt.Scanln(&input2)
	
	// Convert inputs to integers
	a, err1 := strconv.Atoi(input1)
	b, err2 := strconv.Atoi(input2)
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. Please enter valid numbers.")
		return
	}

	// Call functions with the provided numbers
	getAdd(a, b)
	getSub(a, b)
}

// Enter your first num: 12
// Enter your second num: 34
// The addition of 12 and 34 is: 46
// The sub of 12 and 34 is: -22
