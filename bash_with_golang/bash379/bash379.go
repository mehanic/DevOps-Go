package main

import (
	"fmt"
)

func main() {
	// Constants and variable assignments
	const PI = 3.14
	var VAR_A = 10
	var VAR_B = VAR_A
	var VAR_C = VAR_B

	// Print variables
	fmt.Println("Lets print 3 variables:")
	fmt.Println(VAR_A)
	fmt.Println(VAR_B)
	fmt.Println(VAR_C)

	// Demonstrating string interpolation
	fmt.Println("We know this will break:")
	fmt.Println("0. The value of PI is PIabc") // No variable "PIabc" exists, but no error like in bash

	// Printing PI with string interpolation
	fmt.Println("And these will work:")
	fmt.Printf("1. The value of PI is %v\n", PI)
	fmt.Printf("2. The value of PI is %.2f\n", PI) // Using %.2f to format as float with 2 decimals
	fmt.Println("3. The value of PI is", PI)

	// Concatenating strings
	STR_A := "Bob"
	STR_B := "Jane"
	fmt.Printf("%s + %s equals Bob + Jane\n", STR_A, STR_B)

	STR_C := STR_A + " + " + STR_B
	fmt.Printf("%s is the same as Bob + Jane too!\n", STR_C)

	// Adding PI to the new string
	fmt.Printf("%s + %.2f\n", STR_C, PI)

	// Exit the program (optional in Go)
}

