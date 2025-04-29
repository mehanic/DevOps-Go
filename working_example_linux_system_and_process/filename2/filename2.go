package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "programming.txt"

	// Write to the file (similar to 'w' mode in Python)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("I love programming.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	_, err = file.WriteString("I love creating new games.\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Open the file for appending (similar to 'a' mode in Python)
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for appending:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("I also love finding meaning in large datasets.\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	_, err = file.WriteString("I love creating apps that can run in a browser.\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
}

