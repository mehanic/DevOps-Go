package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Give me two numbers, and I'll divide them.")
	fmt.Println("Enter 'q' to quit.")

	for {
		// First number input
		var firstNumber string
		fmt.Print("\nFirst number: ")
		fmt.Scan(&firstNumber)
		if firstNumber == "q" {
			break
		}

		// Second number input
		var secondNumber string
		fmt.Print("Second number: ")
		fmt.Scan(&secondNumber)
		if secondNumber == "q" {
			break
		}

		// Try to convert strings to integers and handle errors
		firstNum, err1 := strconv.Atoi(firstNumber)
		secondNum, err2 := strconv.Atoi(secondNumber)

		if err1 != nil || err2 != nil {
			fmt.Println("Please enter valid numbers.")
			continue
		}

		// Handle division and check for division by zero
		if secondNum == 0 {
			fmt.Println("You can't divide by 0!")
		} else {
			answer := float64(firstNum) / float64(secondNum)
			fmt.Println(answer)
		}
	}
}


// Give me two numbers, and I'll divide them.
// Enter 'q' to quit.

// First number: 23
// Second number: 67
// 0.34328358208955223

// First number: 45
// Second number: 12
// 3.75

// First number: ^Csignal: interrupt
