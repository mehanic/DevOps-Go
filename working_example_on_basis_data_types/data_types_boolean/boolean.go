package main

import (
	"fmt"
)

func main() {
	// Define boolean variables
	a := true
	b := false

	// Print boolean values
	fmt.Println(a)
	fmt.Println(b)

	// Boolean conversion of integers
	fmt.Println(boolFromInt(0))
	fmt.Println(boolFromInt(1))
	fmt.Println(boolFromInt(2))

	// Boolean conversion of strings
	var c string
	fmt.Println(boolFromString(c))
	c = "Some Value"
	fmt.Println(boolFromString(c))
}

// Helper function to convert an integer to a boolean
func boolFromInt(n int) bool {
	return n != 0
}

// Helper function to convert a string to a boolean
func boolFromString(s string) bool {
	return s != ""
}

