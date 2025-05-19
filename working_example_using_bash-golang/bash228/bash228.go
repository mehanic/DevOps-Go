package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Open the user information file (for example, /etc/passwd)
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize a counter for the number of users processed
	cnt := 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Define regular expressions to match the criteria
	neverLoggedInRegex := regexp.MustCompile(`Never logged in`)
	usernameRegex := regexp.MustCompile(`^Username`)
	rootRegex := regexp.MustCompile(`^root`)

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line matches any of the excluded patterns
		if !neverLoggedInRegex.MatchString(line) && 
			!usernameRegex.MatchString(line) && 
			!rootRegex.MatchString(line) {
			
			// Increment the counter and print the line
			cnt++
			fmt.Println(line)
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the total number of users processed
	fmt.Println("========================")
	fmt.Printf("Total Number of Users Processed: %d\n", cnt)
}

