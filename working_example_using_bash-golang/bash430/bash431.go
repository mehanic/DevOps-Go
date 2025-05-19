package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Define the input file path
	myInput := "/home/mansijoshi/Desktop/users.csv"

	// Open the CSV file
	file, err := os.Open(myInput)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	// Skip the header if there is one
	_, err = reader.Read()
	if err != nil {
		fmt.Printf("Failed to read header: %v\n", err)
		return
	}

	// Create slices to hold data
	var aSurname, aName, aUsername, aDepartment, aPassword []string

	// Read each record from the CSV
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file or an error
		}

		if len(record) < 5 {
			continue // Skip incomplete records
		}

		// Append data to slices
		aSurname = append(aSurname, record[0])
		aName = append(aName, record[1])
		aUsername = append(aUsername, record[2])
		aDepartment = append(aDepartment, record[3])
		aPassword = append(aPassword, record[4])
	}

	// Create user accounts
	for index := range aUsername {
		// Create the command to add the user
		passwordHash := hashPassword(aPassword[index])
		cmd := exec.Command("useradd", "-g", aDepartment[index], "-d", "/home/"+aUsername[index],
			"-s", "/bin/bash", "-p", passwordHash, aUsername[index])

		// Run the command
		if err := cmd.Run(); err != nil {
			fmt.Printf("Failed to create user %s: %v\n", aUsername[index], err)
		} else {
			fmt.Printf("User %s created successfully.\n", aUsername[index])
		}
	}
}

// hashPassword hashes the password using openssl
func hashPassword(password string) string {
	// Create a command to hash the password
	cmd := exec.Command("openssl", "passwd", "-1", password)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Failed to hash password: %v\n", err)
		return ""
	}

	// Return the hashed password
	return strings.TrimSpace(string(output))
}

