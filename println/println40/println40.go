package main

import (
	"fmt"
	"strings"
)

// Function to print the picnic items
func printPicnic(itemsMap map[string]int, leftWidth, rightWidth int) {
	// Print the header
	header := "PICNIC ITEMS"
	fmt.Println(center(header, leftWidth+rightWidth, '-'))

	// Print each item in the map
	for k, v := range itemsMap {
		fmt.Println(leftJustify(k, leftWidth) + rightJustify(fmt.Sprintf("%d", v), rightWidth))
	}
}

// Function to center a string within a given width using a specific character
func center(s string, width int, c rune) string {
	padding := width - len(s)
	if padding > 0 {
		leftPadding := padding / 2
		rightPadding := padding - leftPadding
		return strings.Repeat(string(c), leftPadding) + s + strings.Repeat(string(c), rightPadding)
	}
	return s
}

// Function to left justify a string within a given width
func leftJustify(s string, width int) string {
	return fmt.Sprintf("%-*s", width, s)
}

// Function to right justify a string within a given width
func rightJustify(s string, width int) string {
	return fmt.Sprintf("%*s", width, s)
}

func main() {
	picnicItems := map[string]int{
		"sandwiches": 4,
		"apples":     12,
		"cups":       4,
		"cookies":    8000,
	}

	printPicnic(picnicItems, 12, 5)
	printPicnic(picnicItems, 20, 6)
}

