package main

import (
	"fmt"
)

func main() {
	// Declare a constant
	const AUTHOR = "Ganesh Naik"

	// Print the author
	fmt.Println("Author:", AUTHOR)

	// The following line would cause a compile-time error because AUTHOR is a constant and cannot be changed
	// AUTHOR = "John"  // Uncommenting this line will result in an error

	// Trying to reassign a constant is not allowed in Go, so no equivalent for AUTHOR="John" exists.
}

