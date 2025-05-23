package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
//	"strings"
)

func main() {
	// Open the old file
	oldFile, err := os.Open("madLibs_old.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer oldFile.Close()

	// Read the content of the old file
	scanner := bufio.NewScanner(oldFile)
	var oldText string
	for scanner.Scan() {
		oldText += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Define the placeholders
	old := []string{"ADJECTIVE", "NOUN", "VERB", "NOUN"}

	// Create a new file for output
	newFile, err := os.Create("madLibs_new.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer newFile.Close()

	newText := oldText

	// For each placeholder, ask for user input and replace in the text
	for _, item := range old {
		var prompt string
		if item == "ADJECTIVE" || item == "ADVERB" {
			prompt = fmt.Sprintf("Enter an %s: ", item)
		} else {
			prompt = fmt.Sprintf("Enter a %s: ", item)
		}

		fmt.Print(prompt)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		replace := scanner.Text()

		// Use regular expressions to replace the first occurrence of the placeholder
		re := regexp.MustCompile(item)
		newText = re.ReplaceAllString(newText, replace)
	}

	// Write the updated content to the new file
	_, err = newFile.WriteString(newText)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Mad Libs completed and written to madLibs_new.txt")
}

