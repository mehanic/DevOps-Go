package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Open the user information file (e.g., /etc/passwd)
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize a counter for the number of users processed
	cnt := 0

	// Print the header
	fmt.Printf("%8s %11s\n", "Username", "Login date")
	fmt.Println("====================")

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Define regular expressions to match the criteria
	neverLoggedInRegex := regexp.MustCompile(`Never logged in`)
	usernameRegex := regexp.MustCompile(`^Username`)
	rootRegex := regexp.MustCompile(`^root`)

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line) // Split the line into fields

		// Check if the line matches any of the excluded patterns
		if !neverLoggedInRegex.MatchString(line) &&
			!usernameRegex.MatchString(line) &&
			!rootRegex.MatchString(line) {

			// Increment the counter
			cnt++

			// Check the number of fields and format the output accordingly
			if len(fields) == 8 {
				fmt.Printf("%8s %2s %3s %4s\n", fields[0], fields[4], fields[3], fields[7])
			} else {
				fmt.Printf("%8s %2s %3s %4s\n", fields[0], fields[5], fields[4], fields[8])
			}
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the footer
	fmt.Println("====================")
	fmt.Printf("Total Number of Users Processed: %d\n", cnt)
}

