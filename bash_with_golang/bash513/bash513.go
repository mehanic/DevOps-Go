package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

// Function to check if the arguments count matches the expected value.
func checkArguments(expected int, args []string) bool {
	return len(args) == expected
}

// Function to check if the input is an integer.
func checkInteger(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

// Function to generate a random password of the specified length.
func generatePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	password := make([]byte, length)
	for i := range password {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Fatal("Error generating random number:", err)
		}
		password[i] = charset[charIndex.Int64()]
	}
	return string(password)
}

func main() {
	// Check for the correct number of arguments.
	if !checkArguments(2, os.Args) {
		fmt.Println("Incorrect usage! Usage: ./library-password-generator <length>")
		os.Exit(1)
	}

	// Verify that the argument is an integer.
	lengthStr := os.Args[1]
	if !checkInteger(lengthStr) {
		fmt.Println("Argument must be an integer!")
		os.Exit(1)
	}

	// Convert the argument to an integer.
	length, _ := strconv.Atoi(lengthStr)

	// Generate and print the password.
	password := generatePassword(length)
	fmt.Println("Your generated password is:", password)
}

