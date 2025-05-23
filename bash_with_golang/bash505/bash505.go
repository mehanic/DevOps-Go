package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Set the current working directory to the directory of the script.
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Change the working directory.
	err = os.Chdir(dir)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Create a reader to capture user input.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type anything you like: ")

	// Read the input until a newline is encountered.
	userInput, _ := reader.ReadString('\n')

	// Save the user's input to a file. Append the input to the file.
	file, err := os.OpenFile("redirect-to-file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write the user's input to the file.
	if _, err := file.WriteString(userInput); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Your input has been saved.")
}

