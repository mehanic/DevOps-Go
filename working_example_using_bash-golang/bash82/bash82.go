package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Define the input file path
	myInput := "/home/mansijoshi/Desktop/input.csv" // Change the filename if necessary

	// Open the input file
	file, err := os.Open(myInput)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	// Declare slices to hold user information
	var surnames, names, usernames, departments, passwords []string

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Iterate over the records
	for _, record := range records {
		if len(record) < 5 {
			fmt.Println("Invalid record:", record)
			continue
		}
		surnames = append(surnames, record[0])
		names = append(names, record[1])
		usernames = append(usernames, record[2])
		departments = append(departments, record[3])
		passwords = append(passwords, record[4])
	}

	// Create users based on the collected data
	for i, username := range usernames {
		// Create the command to add the user
		cmd := exec.Command("useradd",
			"-g", departments[i],
			"-d", "/home/"+username,
			"-s", "/bin/bash",
			"-p", hashPassword(passwords[i]),
			username,
		)

		// Run the command and check for errors
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to add user %s: %v\n", username, err)
		} else {
			fmt.Printf("User %s added successfully.\n", username)
		}
	}
}

// hashPassword uses OpenSSL to hash the password
func hashPassword(password string) string {
	// Create a command to hash the password using OpenSSL
	cmd := exec.Command("openssl", "passwd", "-1", password)

	// Capture the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return ""
	}

	// Return the trimmed output
	return strings.TrimSpace(string(output))
}

