package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	// Predefined passwords map
	PASSWORDS := map[string]string{
		"email":   "F7minlBDDuvMJuxESSKHFhTxFtjVB6",
		"blog":    "VmALvQyKAxiVH5G8v01if1MLZF3sdt",
		"luggage": "12345",
	}

	// Check if account is passed as an argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [account] - copy account password")
		os.Exit(1)
	}

	// Get the account name from command-line arguments
	account := os.Args[1]

	// Check if the account exists in the PASSWORDS map
	if password, exists := PASSWORDS[account]; exists {
		// Copy password to clipboard
		err := clipboard.WriteAll(password)
		if err != nil {
			fmt.Println("Failed to copy password to clipboard:", err)
			os.Exit(1)
		}
		fmt.Printf("Password for %s copied to clipboard.\n", account)
	} else {
		fmt.Printf("There is no account named %s.\n", account)
	}
}

