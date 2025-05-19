package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get the total number of parameters (excluding the program name)
	paramCount := len(os.Args) - 1
	fmt.Printf("Total number of parameters are = %d\n", paramCount)

	// Get the script name
	scriptName := os.Args[0]
	fmt.Printf("Script name = %s\n", scriptName)

	// Print individual parameters if they exist
	params := os.Args[1:] // Slice of parameters excluding the script name

	if len(params) >= 1 {
		fmt.Printf("First Parameter is %s\n", params[0])
	} else {
		fmt.Println("First Parameter is")
	}

	if len(params) >= 2 {
		fmt.Printf("Second Parameter is %s\n", params[1])
	} else {
		fmt.Println("Second Parameter is")
	}

	if len(params) >= 3 {
		fmt.Printf("Third Parameter is %s\n", params[2])
	} else {
		fmt.Println("Third Parameter is")
	}

	if len(params) >= 4 {
		fmt.Printf("Fourth Parameter is %s\n", params[3])
	} else {
		fmt.Println("Fourth Parameter is")
	}

	if len(params) >= 5 {
		fmt.Printf("Fifth Parameter is %s\n", params[4])
	} else {
		fmt.Println("Fifth Parameter is")
	}

	// Print all parameters
	allParams := strings.Join(params, " ")
	fmt.Printf("All parameters are = %s\n", allParams)
}

