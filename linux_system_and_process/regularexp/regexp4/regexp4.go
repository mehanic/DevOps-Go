package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := `PYTHON2 and there are two major
vers python3 and python in future python4`

	pat := `\bpython[23]?\b`

	// Compile the regular expression with case-insensitive flag
	re, err := regexp.Compile("(?i)" + pat)
	if err != nil {
		fmt.Println("Error compiling regular expression:", err)
		return
	}

	// Search for the first match
	match := re.FindStringSubmatch(text)
	if match != nil {
		fmt.Printf("Match from your pattern: %s\n", match[0])

		// Get the start and end index of the match
		start := re.FindStringIndex(text)[0]
		end := re.FindStringIndex(text)[1] - 1
		length := end - start + 1

		fmt.Printf("Starting index: %d\n", start)
		fmt.Printf("Ending index: %d\n", end)
		fmt.Printf("Length: %d\n", length)
	} else {
		fmt.Println("No match found")
	}
}

