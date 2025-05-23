package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	// Check if the number is even
	result, err := evenNum(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The number you entered: ", result)
	}

	//***************************************

	// Calculate the square root
	result1, err1 := sRoot(-5)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}

	//**********************************************

	// Open a file
	file, err9 := os.Open("test.txt")
	if err9 != nil {
		fmt.Println(err9)
	} else {
		fmt.Println("Our file:", file)
	}
}

// Function to check if a number is even
func evenNum(num int) (int, error) {
	if num%2 != 0 {
		return 0, errors.New("ERROR: You did not enter an even number")
	}
	return num, nil
}

// Function to calculate the square root
func sRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Square root of negative numbers cannot be calculated")
	}
	return math.Sqrt(num), nil
}

// ERROR: You did not enter an even number
// Square root of negative numbers cannot be calculated
// open test.txt: no such file or directory
