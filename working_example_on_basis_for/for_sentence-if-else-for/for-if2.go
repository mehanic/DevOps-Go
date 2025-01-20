package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Initial sentence with placeholders
	text := "The ADJECTIVE panda walked to the NOUN and then VERB. A nearby NOUN was unaffected by those events."

	// Define the placeholders to replace
	placeholders := []string{"ADJECTIVE", "NOUN", "VERB"}

	// For each placeholder, ask the user for a replacement
	for _, placeholder := range placeholders {
		var prompt string
		if placeholder == "ADJECTIVE" || placeholder == "ADVERB" {
			prompt = fmt.Sprintf("Enter an %s: ", placeholder)
		} else {
			prompt = fmt.Sprintf("Enter a %s: ", placeholder)
		}

		// Get user input
		fmt.Print(prompt)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		replacement := scanner.Text()

		// Use regular expressions to replace the first occurrence of the placeholder
		re := regexp.MustCompile(placeholder)
		text = re.ReplaceAllString(text, replacement)
	}

	// Print the new sentence with user inputs
	fmt.Println("\nNew sentence:")
	fmt.Println(text)
}

