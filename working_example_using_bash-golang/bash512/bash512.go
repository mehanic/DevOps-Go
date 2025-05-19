package main

import (
	"fmt"
	"os"
)

// Declare a constant (this is global in scope).
const constantVariable = "constant"

// Declare a global variable to hold the input.
var inputVariable string

// Function that demonstrates variable scoping.
func helloVariable() {
	// Define a local constant inside the function, which shadows the global constant.
	const constantVariable = "maybe not so constant?"

	// Print the values inside the function.
	fmt.Println("This is the input variable:", inputVariable)
	fmt.Println("This is the constant:", constantVariable)
}

func main() {
	// Check if the user supplied at least one argument.
	if len(os.Args) < 2 {
		fmt.Println("Missing an argument!")
		fmt.Println("Usage: ./functions-and-variables <input>")
		os.Exit(1)
	}

	// Assign the first argument to the inputVariable.
	inputVariable = os.Args[1]

	// Call the function that prints variable values.
	helloVariable()

	// Access the constant outside the function.
	fmt.Println("Constant outside function:", constantVariable)
}

