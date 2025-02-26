package main

import (
	"fmt"
	"strings"
)

// Define a function type that takes a string and returns a string
type StringFunc func(string) string

// Decorator function that takes another function and returns a new function
func decorator(functionToDecorate StringFunc) StringFunc {
	return func(value string) string {
		fmt.Printf("you are calling %T with '%s' as parameter\n", functionToDecorate, value)
		return functionToDecorate(value)
	}
}

// Function to replace commas with spaces
func replaceCommasWithSpaces(value string) string {
	return strings.Replace(value, ",", " ", -1)
}

func main() {
	// Apply the decorator
	functionToUse := decorator(replaceCommasWithSpaces)

	// Call the decorated function
	result := functionToUse("my,commas,will,be,replaces,with,spaces")
	fmt.Println(result)
}

// you are calling main.StringFunc with 'my,commas,will,be,replaces,with,spaces' as parameter
// my commas will be replaces with spaces
