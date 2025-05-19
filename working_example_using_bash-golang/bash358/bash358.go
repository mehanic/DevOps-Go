package main

import (
	"fmt"
)

func main() {
	num1 := 10
	num2 := 30

	// Check if num1 is less than num2 and print the result (true or false)
	isLess := num1 < num2
	fmt.Println(isLess)

	// Check if num1 is less than num2 using if statement and print exit code (0 for true, 1 for false)
	if num1 < num2 {
		fmt.Println(0) // Equivalent to exit code 0
	} else {
		fmt.Println(1) // Equivalent to exit code 1
	}
}

