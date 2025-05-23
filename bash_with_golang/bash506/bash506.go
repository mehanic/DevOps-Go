package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const passwordLength = 20 // Desired password length

func main() {
	password := generateRandomPassword(passwordLength)
	fmt.Printf("Your random password is: %s\n", password)
}

// generateRandomPassword generates a random password of a specified length.
func generateRandomPassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := range result {
		// Generate a random index into the charset
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err) // Handle errors in a real application
		}
		result[i] = charset[index.Int64()]
	}

	return string(result)
}

