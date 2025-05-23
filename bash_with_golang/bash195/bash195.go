package main

import (
	"fmt"
	"strings"
)

// printPicnic function prints the picnic items with left and right width formatting.
func printPicnic(itemsDict map[string]int, leftWidth, rightWidth int) {
	// Print the header centered, with dashes filling the sides.
	fmt.Println(center("PICNIC ITEMS", leftWidth+rightWidth, '-'))

	// Print each key-value pair with formatting.
	for k, v := range itemsDict {
		// Left-justify the key and right-justify the value
		fmt.Printf("%-*s%*d\n", leftWidth, k, rightWidth, v)
	}
}

// center function centers the text between two widths with a specific padding character.
func center(text string, width int, padChar rune) string {
	if len(text) >= width {
		return text
	}
	padLen := (width - len(text)) / 2
	return strings.Repeat(string(padChar), padLen) + text + strings.Repeat(string(padChar), width-len(text)-padLen)
}

func main() {
	// Create the items dictionary
	picnicItems := map[string]int{
		"sandwiches": 4,
		"apples":     12,
		"cups":       4,
		"cookies":    8000,
	}

	// Print picnic items with different widths
	printPicnic(picnicItems, 12, 5)
	printPicnic(picnicItems, 20, 6)
}

