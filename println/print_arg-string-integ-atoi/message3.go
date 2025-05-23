package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstStr, secondStr, thirdStr string
	var first, second, third int
	var err error

	// Prompt user and read input for first number
	fmt.Print("Enter the first number to add: ")
	_, err = fmt.Scan(&firstStr)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Prompt user and read input for second number
	fmt.Print("Enter the second number to add: ")
	_, err = fmt.Scan(&secondStr)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Prompt user and read input for third number
	fmt.Print("Enter the third number to add: ")
	_, err = fmt.Scan(&thirdStr)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Convert input strings to integers
	first, err = strconv.Atoi(firstStr)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	second, err = strconv.Atoi(secondStr)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	third, err = strconv.Atoi(thirdStr)
	if err != nil {
		fmt.Println("Error converting input to integer:", err)
		return
	}

	// Calculate the sum
	sum := first + second + third

	// Print the result
	fmt.Printf("The sum is %d\n", sum)
}

// Enter the first number to add: 45
// Enter the second number to add: 56
// Enter the third number to add: 34
// The sum is 135
