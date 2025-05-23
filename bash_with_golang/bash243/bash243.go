package main

import (
	"fmt"
	"os"
)

// Define color escape codes as constants
const (
	RED    = "\033[31m"
	GREEN  = "\033[32m"
	BLUE   = "\033[34m"
	RESET  = "\033[0m"
)

func main() {
	// Check if at least one argument is provided
	if len(os.Args) < 2 {
		// Print usage message in red
		fmt.Printf("%sUsage: %s <name>%s\n", RED, os.Args[0], RESET)
		os.Exit(1)
	}

	// Print greeting message in green
	fmt.Printf("%sHello %s%s\n", GREEN, os.Args[1], RESET)
	os.Exit(0)

	// Extra check (not necessary due to the above check)
	if os.Args[1] == "" {
		fmt.Printf("Usage: %s <name>\n", os.Args[0])
		os.Exit(2)
	}
}

