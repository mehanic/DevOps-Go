package main

import (
	"fmt"
)

func main() {
	// Variable declarations
	USER_AGE := 18
	AGE_LIMIT := 18
	NAME := "Bob" // Change to your username if you want to execute the nested logic
	HAS_NIGHTMARES := true
	USER := "Bob" // Change this value to test different logic

	// Conditional logic
	if USER == NAME {
		if USER_AGE >= AGE_LIMIT {
			if HAS_NIGHTMARES {
				fmt.Printf("%s gets nightmares, and should not see the movie\n", USER)
			}
		}
	} else {
		fmt.Println("Who is this?")
	}
}

