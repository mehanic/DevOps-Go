package main

import (
	"fmt"
	"math/rand"
//	"os"
	"time"
)

const passwordLength = 20

// chars contains the characters to use for password generation.
const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Generate a random password.
	password := generatePassword(passwordLength)

	// Print the generated password.
	fmt.Printf("Your random password is: %s\n", password)
}

// generatePassword generates a random password of the specified length.
func generatePassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		password[i] = chars[rand.Intn(len(chars))] // Select a random character.
	}
	return string(password)
}

