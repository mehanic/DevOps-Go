package main

import (
	"fmt"
	"strconv"
)

// Function to calculate and print the result
func getResult(value int) {
	result := value + 10
	fmt.Printf("Your result is: %d\n", result)
}

func main() {
	// Prompt user for input
	fmt.Print("Enter your number: ")
	var input string
	fmt.Scanln(&input)

	// Convert input to integer
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	// Call function with the provided number
	getResult(num)
}



// Enter your number: 34
// Your result is: 44
