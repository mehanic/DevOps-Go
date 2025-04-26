package main

import (
	"fmt"
	"strings"
)

// Define a function type that takes a string and returns a string
type StringFunc func(string) string

// Decorator function that takes another function and returns a new function
func decorator(functionName string, functionToDecorate StringFunc) StringFunc {
	return func(value string) string {
		// Print the function's name and parameter
		fmt.Printf("you are calling %s with '%s' as parameter\n", functionName, value)
		return functionToDecorate(value)
	}
}

// Function to replace commas with spaces
func replaceCommasWithSpaces(value string) string {
	return strings.Replace(value, ",", " ", -1)
}

func main() {
	// Applying the decorator to the function, and manually passing the function name
	decoratedFunc := decorator("replaceCommasWithSpaces", replaceCommasWithSpaces)

	// Manually setting "name", "module", and "docstring"
	funcName := "replaceCommasWithSpaces"
	funcModule := "main"
	funcDoc := "replaceCommasWithSpaces replaces all commas in a string with spaces."

	// Print name, module, and docstring manually
	fmt.Println("Function Name:", funcName)
	fmt.Println("Function Module:", funcModule)
	fmt.Println("Function Doc:", funcDoc)

	// Calling the decorated function
	result := decoratedFunc("my,commas,will,be,replaces,with,spaces")
	fmt.Println(result)
}

// Function Name: replaceCommasWithSpaces
// Function Module: main
// Function Doc: replaceCommasWithSpaces replaces all commas in a string with spaces.
// you are calling replaceCommasWithSpaces with 'my,commas,will,be,replaces,with,spaces' as parameter
// my commas will be replaces with spaces
