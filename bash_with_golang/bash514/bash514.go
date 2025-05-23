package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Function to set the current working directory to the script's location.
func setCwd() {
	exePath, err := os.Executable() // Get the path of the current executable
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		os.Exit(1)
	}
	dir := filepath.Dir(exePath)
	if err := os.Chdir(dir); err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}
}

// Function to get user input from the console.
func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return input
}

// Function to append the user's input to a file.
func appendToFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}

func main() {
	// Set current working directory to script location.
	setCwd()

	// Capture user input.
	userInput := getUserInput("Type anything you like: ")

	// Append user input to the file.
	appendToFile("redirect-to-file.txt", userInput)

	fmt.Println("User input has been written to redirect-to-file.txt")
}

