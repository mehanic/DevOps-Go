package main

import (
	"fmt"
	"os"
)

// printColored prints a string in the specified color using ANSI escape codes.
func printColored(text string, color string) error {
	colorCode := ""

	// Define ANSI color codes based on the input.
	switch color {
	case "red":
		colorCode = "\033[31m"
	case "blue":
		colorCode = "\033[34m"
	case "green":
		colorCode = "\033[32m"
	default:
		colorCode = "\033[39m" // Default terminal color
	}

	// Print the colored text and reset color at the end.
	fmt.Printf("%s%s\033[39m\n", colorCode, text)
	return nil
}

func main() {
	// Example usage of the printColored function.
	if err := printColored("Hello world!", "red"); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := printColored("Hello world!", "blue"); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := printColored("Hello world!", "green"); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if err := printColored("Hello world!", "magenta"); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

