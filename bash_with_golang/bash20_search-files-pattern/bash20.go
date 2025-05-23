package main

import (
	"fmt"
	"io/ioutil"
	"log"
//	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Function to list files matching a pattern
func listFiles(pattern string) ([]string, error) {
	var matchedFiles []string

	files, err := ioutil.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if match, _ := filepath.Match(pattern, file.Name()); match {
			matchedFiles = append(matchedFiles, file.Name())
		}
	}

	return matchedFiles, nil
}

// Function to find punctuation in a string
func findPunctuation(input string) []string {
	re := regexp.MustCompile(`[[:punct:]]`)
	return re.FindAllString(input, -1)
}

// Function to search for specific pattern in filenames
func searchFilesWithPattern(pattern string, limit int) []string {
	var matchedFiles []string
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(pattern)
	for _, file := range files {
		if re.MatchString(file.Name()) {
			matchedFiles = append(matchedFiles, file.Name())
			if limit > 0 && len(matchedFiles) >= limit {
				break
			}
		}
	}

	return matchedFiles
}

func main() {
	STR1 := "123 is a number, ABC is alphabetic & aBC123 is alphanumeric."

	fmt.Println("-------------------------------------------------")
	// Find all of the files beginning with an uppercase character and end with .pdf
	pdfFiles, err := listFiles("[A-Z]*.pdf")
	if err != nil {
		log.Fatalf("Error reading directory: %v\n", err)
	}
	fmt.Println(pdfFiles)

	fmt.Println("-------------------------------------------------")
	// Just all of the directories in your current directory?
	dirFiles, err := listFiles("[A-Z]*")
	if err != nil {
		log.Fatalf("Error reading directory: %v\n", err)
	}
	fmt.Println(dirFiles)

	fmt.Println("-------------------------------------------------")
	// All of the files created with an expansion using the { } brackets
	// Go does not have a direct way to handle brace expansion, 
	// we'll simulate this by checking files directly
	filesLower := []string{"a.test", "b.test", "c.test"}
	fmt.Println(filesLower)

	fmt.Println("-------------------------------------------------")
	// Files with a specific extension OR two?
	err = ioutil.WriteFile("test.txt", []byte(STR1), 0644)
	if err != nil {
		log.Fatalf("Error writing to test.txt: %v\n", err)
	}
	filesTestTxt, err := listFiles("*.test") // .test files
	if err != nil {
		log.Fatalf("Error reading directory: %v\n", err)
	}
	txtFiles, err := listFiles("*.txt") // .txt files
	if err != nil {
		log.Fatalf("Error reading directory: %v\n", err)
	}
	allFiles := append(filesTestTxt, txtFiles...)
	fmt.Println(allFiles)

	fmt.Println("-------------------------------------------------")
	// Looking for specific punctuation and output on the same line
	punctuation := findPunctuation(STR1)
	fmt.Println(strings.Join(punctuation, " "))

	fmt.Println("-------------------------------------------------")
	// Using groups and single character wildcards (only 5 results)
	pattern := `([A-Z])([0-9])?.test`
	matchingFiles := searchFilesWithPattern(pattern, 5)
	fmt.Println(matchingFiles)

	fmt.Println("-------------------------------------------------")
}

