package main

import (
	"fmt"
	"strings"
)

func main() {
	// Build a long string of equals signs
	divider := strings.Repeat("=", 29) // 29 equals signs
	divider += divider                  // Concatenate two dividers

	// Format strings for output
	header := "\n %-10s %11s %8s %10s\n"
	format := " %-10s %11.2f %8d %10.2f\n"

	// Width of divider
	totalWidth := 44

	// Print categories
	fmt.Printf(header, "ITEM", "PER UNIT", "NUM", "TOTAL")

	// Print divider to match width of report
	fmt.Printf("%*s\n", totalWidth, divider)

	// Print lines of report
	fmt.Printf(format,
		"Chair", 79.95, 4, 319.80,
		"Table", 209.99, 1, 209.99,
		"Armchair", 315.49, 2, 630.98,
	)
}

