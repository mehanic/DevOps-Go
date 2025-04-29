package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Input text
	text := "is a python learn and \n it is easy \tto"

	// Define the regex pattern (modify which pattern is active by changing the variable)
	// Uncomment one of the patterns to match different parts of the text
	 myPat := `^i[ts]`     // Matches strings starting with 'i' followed by 't' or 's'
	// myPat := `learn$`      // Matches the word 'learn' at the end of the string
	// myPat := `\blearn\b`   // Matches the word 'learn' as a whole word (word boundary)
//	 myPat := `\Blearn\B`   // Matches 'learn' when not at a word boundary
	//myPat := `\n`           // Matches newline character

	// Compile the regular expression
	re := regexp.MustCompile(myPat)

	// Find all occurrences of the pattern
	matches := re.FindAllString(text, -1)

	// Print the matches
	fmt.Println(matches)
}

