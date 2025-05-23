package main

import (
	"fmt"
	"io/ioutil"
	"log"
//	"os"
//	"path/filepath"
	"regexp"
	"strings"
)

// Function to list files matching a pattern
func listFilesMatchingPattern(pattern string) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("Failed to read directory: %s", err)
	}

	for _, file := range files {
		if match, _ := regexp.MatchString(pattern, file.Name()); match {
			fmt.Println(file.Name())
		}
	}
}

// Function to filter files by a specific extension
func listFilesByExtension(extensions []string) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("Failed to read directory: %s", err)
	}

	for _, file := range files {
		for _, ext := range extensions {
			if strings.HasSuffix(file.Name(), ext) {
				fmt.Println(file.Name())
			}
		}
	}
}

// Function to find punctuation in a string
func findPunctuation(str string) {
	re := regexp.MustCompile(`[[:punct:]]`)
	punctuations := re.FindAllString(str, -1)
	fmt.Println(strings.Join(punctuations, " "))
}

// Function to demonstrate string matching and tail functionality
func matchUpperDigitFileNames() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalf("Failed to read directory: %s", err)
	}

	var matchedFiles []string
	re := regexp.MustCompile(`([A-Z])([0-9]?)\.test?`)

	for _, file := range files {
		if match := re.MatchString(file.Name()); match {
			matchedFiles = append(matchedFiles, file.Name())
		}
	}

	// Print the last 5 matched results
	if len(matchedFiles) > 5 {
		matchedFiles = matchedFiles[len(matchedFiles)-5:]
	}
	fmt.Println(strings.Join(matchedFiles, "\n"))
}

// Main function
func main() {
	// Example string for testing
	STR1 := "123 is a number, ABC is alphabetic & aBC123 is alphanumeric."

	fmt.Println("-------------------------------------------------")
	// Find all of the files beginning with an uppercase character and end with .pdf
	fmt.Println("Files beginning with an uppercase character and ending with .pdf:")
	listFilesMatchingPattern(`^[A-Z].*\.pdf$`)

	fmt.Println("-------------------------------------------------")
	// Just all of the directories in your current directory
	fmt.Println("Directories starting with an uppercase character:")
	listFilesMatchingPattern(`^[A-Z].*`)

	fmt.Println("-------------------------------------------------")
	// Files with a specific extension OR two
	fmt.Println("Files with .test or .txt extensions:")
	// Create the test.txt file for demonstration purposes
	ioutil.WriteFile("test.txt", []byte(STR1), 0644)
	listFilesByExtension([]string{".test", ".txt"})

	fmt.Println("-------------------------------------------------")
	// Looking for specific punctuation
	fmt.Println("Punctuation found in the string:")
	findPunctuation(STR1)

	fmt.Println("-------------------------------------------------")
	// Using groups and single character wildcards
	fmt.Println("Files matching specific regex pattern (last 5 results):")
	matchUpperDigitFileNames()

	fmt.Println("-------------------------------------------------")
	fmt.Println("Finished processing.")
}

