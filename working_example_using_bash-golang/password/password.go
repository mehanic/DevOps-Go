package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"syscall"
)

func main() {
	// Prompt the user for a password
	fmt.Print("password: ")

	// Read password input without echoing it
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("\nError reading password:", err)
		return
	}

	// Print a newline after password input (since it's suppressed in ReadPassword)
	fmt.Println()

	// Convert the password from byte slice to string and print it (for demonstration purposes)
	// In real-world usage, you wouldn't print the password
	passStr := string(password)
	fmt.Println("Password entered:", passStr)
}

