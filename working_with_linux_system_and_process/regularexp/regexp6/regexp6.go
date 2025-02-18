package main

import (
	"fmt"
	"regexp"
)

func main() {
	myStr := "This is about python. Python is easy to learn and we have two major versions: python2 and python3 "

	// Compile the regular expression
	myPat := `\bPython[23]?\b`
	re, err := regexp.Compile("(?i)" + myPat) // Compile with case-insensitive flag
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Print the compiled regex object
	fmt.Println("Compiled regex:", re)

	// Search for the first match
	match := re.FindString(myStr)
	fmt.Println("First match:", match)

	// Find all matches
	matches := re.FindAllString(myStr, -1)
	fmt.Println("All matches:", matches)

	// Split the string by the pattern
	split := re.Split(myStr, -1)
	fmt.Println("Split result:", split)
}

