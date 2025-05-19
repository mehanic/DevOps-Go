package main

import (
	"fmt"
)

func main() {
	// Variable declaration
	AGE := 21

	// Conditional logic
	if AGE < 18 {
		fmt.Println("You must be 18 or older to see this movie")
	} else if AGE == 21 {
		fmt.Println("You may see the movie and get popcorn")
	} else {
		fmt.Println("You may see the movie!")
		return // Equivalent to `exit 1` in Go
	}

	// This line might not get executed if the return statement is hit
	fmt.Println("This line might not get executed")
}

