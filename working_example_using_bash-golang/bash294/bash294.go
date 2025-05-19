package main

import (
	"fmt"
)

// Declare a global variable
var VAR = "global variable"

// Function that uses a local variable with the same name as the global one
func bash() {
	// Declare a local variable inside the function
	VAR := "local variable"
	fmt.Println(VAR)
}

func main() {
	// Print the global variable
	fmt.Println(VAR)

	// Call the bash function which uses the local variable
	bash()

	// Print the global variable again to show it hasn't changed
	fmt.Println(VAR)
}

