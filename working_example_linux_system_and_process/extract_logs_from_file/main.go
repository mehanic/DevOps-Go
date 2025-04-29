package main

import (
	"fmt"
	"os"
	"log"
	"regexp"
)

// Define a function to extract email addresses from a log file
func extractEmailsFromLog(logFilePath string) []string {
	// Read the log file content
	data, err := os.ReadFile(logFilePath)
	if err != nil {
		log.Fatalf("Failed to read the log file: %s", err)
	}

	// Regular expression to match email addresses
	re := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}\b`)

	// Find all matches in the log data
	emailIDs := re.FindAllString(string(data), -1)

	return emailIDs
}

func main() {
	// Call the function and print the extracted email addresses
	emails := extractEmailsFromLog("sample_email_log.txt")
	fmt.Println("Extracted Email Addresses:")
	for _, email := range emails {
		fmt.Println(email)
	}
}


// go run main.go < sample_email_log.txt                                                                                                                        
// Extracted Email Addresses:
// john.doe@example.com
// support@company.com
// feedback@service.org