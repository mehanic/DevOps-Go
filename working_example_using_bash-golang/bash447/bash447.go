package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func askOption(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s (y/n): ", prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "y" || input == "yes" {
			return true
		} else if input == "n" || input == "no" {
			return false
		} else {
			fmt.Println("Please enter 'y' or 'n'")
		}
	}
}

func main() {
	fmt.Println("Enable the account options you want:")

	homeDir := askOption("1) Home directory (default: on)")
	signatureFile := askOption("2) Signature file (default: off)")
	simplePassword := askOption("3) Simple password (default: off)")

	// Output the selected options
	fmt.Println("\nSelected options:")
	if homeDir {
		fmt.Println("1) Home directory: enabled")
	} else {
		fmt.Println("1) Home directory: disabled")
	}

	if signatureFile {
		fmt.Println("2) Signature file: enabled")
	} else {
		fmt.Println("2) Signature file: disabled")
	}

	if simplePassword {
		fmt.Println("3) Simple password: enabled")
	} else {
		fmt.Println("3) Simple password: disabled")
	}
}

