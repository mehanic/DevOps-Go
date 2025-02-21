package main

import (
	"fmt"
	"strings"
)

func main() {
	// String formatting examples
	name := "ada lovelace"
	fmt.Println(strings.Title(name))  // Equivalent of Python's title()
	fmt.Println(strings.ToUpper(name)) // Uppercase
	fmt.Println(strings.ToLower(name)) // Lowercase

	// String concatenation
	firstName := "ada"
	lastName := "lovelace"
	fullName := firstName + " " + lastName
	fmt.Println(fullName)

	// Formatting and output
	message := "Hello, " + strings.Title(fullName) + "!"
	fmt.Println(message)

	// Tab and newline characters
	fmt.Println("\tPython")
	fmt.Println("Languages:\n\tPython\n\tC\n\tJavaScript")

	// Define a favorite language variable
	favoriteLanguage := "Python "
	_ = favoriteLanguage // This prevents an unused variable error
}

// Ada Lovelace
// ADA LOVELACE
// ada lovelace
// ada lovelace
// Hello, Ada Lovelace!
// 	Python
// Languages:
// 	Python
// 	C
// 	JavaScript
