package main

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// Get text from clipboard
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Failed to read from clipboard:", err)
		return
	}

	// Separate lines and add stars
	lines := strings.Split(text, "\n")
	for i := range lines {
		lines[i] = "* " + lines[i]
	}
	text = strings.Join(lines, "\n")

	// Copy the modified text to clipboard
	err = clipboard.WriteAll(text)
	if err != nil {
		fmt.Println("Failed to write to clipboard:", err)
		return
	}

	fmt.Println("Text modified and copied to clipboard successfully.")
}

//Text modified and copied to clipboard successfully. (copied before)
