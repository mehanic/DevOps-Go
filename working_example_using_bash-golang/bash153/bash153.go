package main

import (
	"fmt"
	"os"
)

var (
	GLOBAL_VAR1 = "one" // Global variable 1
	GLOBAL_VAR2 = "two" // Global variable 2
)

// functionOne is a function that simulates the Bash function.
func functionOne() {
	LOCAL_VAR1 := "one" // Local variable
	// <Replace with function code.>
	fmt.Println("Inside functionOne, LOCAL_VAR1:", LOCAL_VAR1)
}

func main() {
	// Main body of the Go program starts here.

	// Call the function
	functionOne()

	// Example of using global variables
	fmt.Println("Global variables:", GLOBAL_VAR1, GLOBAL_VAR2)

	// Exit with an explicit exit status
	os.Exit(0)
}

