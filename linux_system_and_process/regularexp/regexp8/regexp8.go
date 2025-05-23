package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

// Define regular expressions for phone numbers and emails
var (
	phoneRegex = regexp.MustCompile(`(?x)
		(\d{3}|\(\d{3}\))?              # area code
		[\s\-\.]?                        # separator
		(\d{3})                          # first 3 digits
		[\s\-\.]                         # separator
		(\d{4})                          # last 4 digits
		[\s]*(?:ext|x|ext\.)[\s]*(\d{2,5})?  # extension
	`)

	emailRegex = regexp.MustCompile(`(?x)
		[a-zA-Z0-9._%+-]+                # username
		@                                # @ symbol
		[a-zA-Z0-9.-]+                   # domain name
		(?:\.[a-zA-Z]{2,4}){1,2}         # dot-something
	`)
)

func main() {
	// Get text from the clipboard
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Error reading from clipboard:", err)
		return
	}

	var matches []string

	// Find phone numbers
	phoneMatches := phoneRegex.FindAllStringSubmatch(text, -1)
	for _, groups := range phoneMatches {
		phoneNum := strings.Join([]string{groups[1], groups[2], groups[3]}, "-")
		if groups[4] != "" {
			phoneNum += " x" + groups[4]
		}
		matches = append(matches, phoneNum)
	}

	// Find email addresses
	emailMatches := emailRegex.FindAllStringSubmatch(text, -1)
	for _, groups := range emailMatches {
		matches = append(matches, groups[0])
	}

	// Copy results to the clipboard
	if len(matches) > 0 {
		result := strings.Join(matches, "\n")
		if err := clipboard.WriteAll(result); err != nil {
			fmt.Println("Error writing to clipboard:", err)
			return
		}
		fmt.Println("Copied to clipboard:")
		fmt.Println(result)
	} else {
		fmt.Println("No phone numbers or email addresses found.")
	}
}

