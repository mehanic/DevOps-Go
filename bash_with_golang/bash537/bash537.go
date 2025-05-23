package main

import (
	"fmt"
)

func main() {
	// Define the strings
	S1 := "HelloWorld"
	S2 := "HelloBash"

	// Check if strings are equal
	if S1 == S2 {
		fmt.Println("Both strings are equal")
	} else {
		fmt.Println("Strings are NOT equal")
	}

	// Check if both strings are not empty
	if len(S1) > 0 && len(S2) > 0 {
		fmt.Println("Both strings are not empty.")
	}
}

