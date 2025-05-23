package main

import (
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	// Passwords map
	passwords := map[string]string{
		"email":   "F7minlBDDuvMJuxESSKHFhTxFtjVB6",
		"blog":    "VmALvQyKAxiVH5G8v01if1MLZF3sdt",
		"luggage": "12345",
	}

	// Check if we have at least one command line argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run pw.go [account] - copy account password")
		return
	}

	// Get the account name from the command line arguments
	account := os.Args[1]

	// Check if the account exists in the map
	if password, exists := passwords[account]; exists {
		// Copy the password to the clipboard
		err := clipboard.WriteAll(password)
		if err != nil {
			fmt.Println("Error copying to clipboard:", err)
			return
		}
		fmt.Println("Password for", account, "copied to clipboard.")
	} else {
		fmt.Println("There is no account named", account)
	}
}

