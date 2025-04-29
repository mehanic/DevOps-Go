package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Define the string and pattern
	myStr := "This is python and we are having python2 and python3 version"
	myPat := `\bpython[23]?\b`

	// Compile the regular expression
	re, err := regexp.Compile(myPat)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	// Find all matches
	matches := re.FindAllStringSubmatchIndex(myStr, -1)

	// Iterate over matches and print details
	for _, match := range matches {
		startIndex := match[0]
		endIndex := match[1]
		matchString := myStr[startIndex:endIndex]
		fmt.Printf("The match is: %s starting index: %d, ending index %d\n", matchString, startIndex, endIndex-1)
	}
}

