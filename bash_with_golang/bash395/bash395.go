package main

import (
	"fmt"
	"io/ioutil"
	"log"
//	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Function to search for a term in files of a given directory
func searchInFiles(directory, searchTerm string, resultFile string) {
	// Get all files in the directory recursively
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatalf("Failed to read directory: %s, error: %s", directory, err)
	}

	var matches []string
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(directory, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				log.Printf("Failed to read file: %s, error: %s", filePath, err)
				continue
			}
			if strings.Contains(string(content), searchTerm) {
				matches = append(matches, filePath)
			}
		}
	}

	// Write results to a result file
	if err := ioutil.WriteFile(resultFile, []byte(strings.Join(matches, "\n")), 0644); err != nil {
		log.Fatalf("Failed to write results to file: %s, error: %s", resultFile, err)
	}
}

// Function to search recursively for terms
func recursiveSearch(directory string, searchTerms []string, resultFile string) {
	cmd := exec.Command("grep", "-r", "-e", strings.Join(searchTerms, " -e "), directory)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error during grep: %s", err)
	}
	if err := ioutil.WriteFile(resultFile, output, 0644); err != nil {
		log.Fatalf("Failed to write results to file: %s, error: %s", resultFile, err)
	}
}

// Function to search for terms in specific file types
func findFilesWithTerms(directory string, searchTerm string, fileType string, resultFile string) {
	cmd := exec.Command("find", directory, "-type", "f", "-name", "*."+fileType, "-print")
	output, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error during find: %s", err)
	}

	// Convert output to a list of files
	files := strings.Split(string(output), "\n")
	var matches []string

	for _, file := range files {
		if file != "" {
			content, err := ioutil.ReadFile(file)
			if err != nil {
				log.Printf("Failed to read file: %s, error: %s", file, err)
				continue
			}
			if strings.Contains(string(content), searchTerm) {
				matches = append(matches, file)
			}
		}
	}

	// Write results to a result file
	if err := ioutil.WriteFile(resultFile, []byte(strings.Join(matches, "\n")), 0644); err != nil {
		log.Fatalf("Failed to write results to file: %s, error: %s", resultFile, err)
	}
}

// Main function
func main() {
	directory := "www.packtpub.com"
	searchTerm := "Packt"

	// Search for the term in all files and write to result1.txt
	searchInFiles(directory, searchTerm, "result1.txt")

	// Perform recursive search for the term
	recursiveSearch(directory, []string{searchTerm}, "result2.txt")

	// Search for multiple terms
	recursiveSearch(directory, []string{searchTerm, "Publishing"}, "result3.txt")

	// Search using find
	findFilesWithTerms(directory, searchTerm, "xml", "result4.txt")

	// Search for specific file types excluding certain patterns
	findFilesWithTerms(directory, searchTerm, "xml", "result5.txt")

	// Wildcard search simulation
	// Note: Go doesn't have direct wildcard handling like Bash, 
	// this is a simple example for demonstration
	searchInFiles(directory, searchTerm, "result6.txt")

	// Check for results
	result, err := ioutil.ReadFile("result1.txt")
	if err != nil {
		log.Fatalf("Failed to read results: %s", err)
	}
	if len(result) > 0 {
		fmt.Println("We found results!")
	} else {
		fmt.Println("It broke - it shouldn't happen (Packt is everywhere)!")
	}

	// Simulating the history command
	cmd := exec.Command("bash", "-c", "history | grep 'ls'")
	historyOutput, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error retrieving history: %s", err)
	}
	fmt.Println(string(historyOutput))

	// Final message
	fmt.Println("We can do a lot with grep!")
}

