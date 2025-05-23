package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileName := "/tmp/grep-then-else.txt"

	// Create the file; it will create the file if it does not exist.
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating/opening the file:", err)
		return
	}
	defer file.Close()

	// Read the content of the file
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Check if the file contains the keyword
	if strings.Contains(string(content), "keyword") {
		// If the file contains the keyword, remove the file.
		if err := os.Remove(fileName); err != nil {
			fmt.Println("Error removing the file:", err)
		}
	} else {
		// Otherwise, write the keyword to the file.
		if _, err := file.WriteString("keyword\n"); err != nil {
			fmt.Println("Error writing to the file:", err)
		}
	}
}

