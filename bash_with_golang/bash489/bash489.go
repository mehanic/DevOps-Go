package main

import (
	"fmt"
	"io/ioutil"
	"os"
//	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Create a directory to store log files with errors.
	errorDirectory := "/tmp/error_logfiles/"
	err := os.MkdirAll(errorDirectory, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Create a list of log files.
	logFiles, err := filepath.Glob("/var/log/*.log")
	if err != nil {
		fmt.Printf("Error finding log files: %v\n", err)
		return
	}

	for _, file := range logFiles {
		// Check if the file contains the word 'error' (case insensitive)
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			continue
		}

		// Use regex to check for 'error' or 'Error'
		if matched, _ := regexp.MatchString(`(?i)error`, string(content)); matched {
			fmt.Printf("%s contains error(s), copying it to archive %s.\n", file, errorDirectory)

			// Copy the file to the error directory
			destFile := filepath.Join(errorDirectory, filepath.Base(file))
			err = ioutil.WriteFile(destFile, content, os.ModePerm)
			if err != nil {
				fmt.Printf("Error copying file to %s: %v\n", destFile, err)
				continue
			}

			// In-place edit, only print lines matching 'error' or 'Error'
			filteredLines := filterErrorLines(string(content))
			err = ioutil.WriteFile(destFile, []byte(filteredLines), os.ModePerm)
			if err != nil {
				fmt.Printf("Error writing filtered content to %s: %v\n", destFile, err)
			}
		}
	}
}

// filterErrorLines filters lines that contain 'error' or 'Error'
func filterErrorLines(content string) string {
	var filteredLines []string
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if matched, _ := regexp.MatchString(`(?i)error`, line); matched {
			filteredLines = append(filteredLines, line)
		}
	}
	return strings.Join(filteredLines, "\n")
}

