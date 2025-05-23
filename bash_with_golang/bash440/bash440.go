package main

import (
	"fmt"
	"os"
	"regexp"
)

func isValidEmail(email string) bool {
	// Define the regex pattern for validating the email
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go <EMAIL>")
		return
	}

	email := os.Args[1]

	if isValidEmail(email) {
		fmt.Printf("%s is valid\n", email)
	} else {
		fmt.Printf("%s is NOT valid\n", email)
	}
}

