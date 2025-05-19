package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	// Check for the correct number of arguments.
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Usage: go run password-generator.go <length>")
		os.Exit(1)
	}

	// Verify the length argument.
	passwordLength, err := strconv.Atoi(os.Args[1])
	if err != nil || passwordLength <= 0 {
		fmt.Println("Please enter a positive length (number).")
		os.Exit(1)
	}

	// Generate the password.
	password, err := generatePassword(passwordLength)
	if err != nil {
		fmt.Println("Error generating password:", err)
		os.Exit(1)
	}

	// Print the generated password.
	fmt.Println(password)
}

// generatePassword creates a random password of the specified length.
func generatePassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		// Generate a random index for the letterBytes string.
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}
		password[i] = letterBytes[index.Int64()]
	}
	return string(password), nil
}

