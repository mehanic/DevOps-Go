package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `this is a string EnD
this is second line enD
This is third line end
asfasd this end`

	// Compile the regular expression
	myPat := `end$`
	re, err := regexp.Compile(myPat)
	if err != nil {
		fmt.Println("Error compiling regular expression:", err)
		return
	}

	// Find all matches with multiline and case-insensitive flags
	matches := re.FindAllString(text, -1)
	fmt.Println(matches)
}

