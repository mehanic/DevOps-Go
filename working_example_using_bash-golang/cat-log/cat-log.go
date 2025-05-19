package main

import (
	"fmt"
	"os"
)

func main() {
	// Define the log content
	logContent := `LOG FILE HEADER
This is a test log file
Function: System statistics
`

	// Create or truncate the log.txt file
	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed at the end of the function

	// Write the log content to the file
	_, err = file.WriteString(logContent)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Log file created successfully.")
}

