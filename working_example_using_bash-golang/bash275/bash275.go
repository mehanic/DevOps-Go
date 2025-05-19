package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the logfile names to search.")
		return
	}

	// Create or open the diagnostics log file
	diagFile, err := os.OpenFile("Diagnostics.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening diagnostics log file:", err)
		return
	}
	defer diagFile.Close()

	for _, logName := range os.Args[1:] {
		if logName == "" {
			fmt.Println("Please provide a valid logfile name to search.")
			continue
		}

		// Call the main log analysis functions
		stuck(logName, diagFile)
		deadlock(logName, diagFile)
		uncheckedException(logName, diagFile)
	}

	// Remove duplicate lines from Diagnostics.log
	removeDuplicates("Diagnostics.log")
}

// Function to find stuck threads
func stuck(logName string, diagFile *os.File) {
	file, err := os.Open(logName)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if containsStuckThread(line) {
			// Write the matched line and the next 4 lines to the diagnostics file
			writeToDiagnostics(scanner, diagFile, line)
		}
	}
}

// Function to find deadlocks
func deadlock(logName string, diagFile *os.File) {
	file, err := os.Open(logName)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if containsDeadlock(line) {
			// Write the matched line and the next 4 lines to the diagnostics file
			writeToDiagnostics(scanner, diagFile, line)
		}
	}
}

// Function to find unchecked exceptions
func uncheckedException(logName string, diagFile *os.File) {
	file, err := os.Open(logName)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if containsUncheckedException(line) {
			// Write the matched line and the next 4 lines to the diagnostics file
			writeToDiagnostics(scanner, diagFile, line)
		}
	}
}

// Function to write to diagnostics log with 4 additional lines
func writeToDiagnostics(scanner *bufio.Scanner, diagFile *os.File, line string) {
	diagFile.WriteString(line + "\n")
	for i := 0; i < 4 && scanner.Scan(); i++ {
		diagFile.WriteString(scanner.Text() + "\n")
	}
}

// Function to check if the line contains a stuck thread
func containsStuckThread(line string) bool {
	pattern := `WL-000337|BEA-000337|WL-101020|WL-101017|WL-000802|WL-101020|BEA-101017`
	matched, _ := regexp.MatchString(pattern, line)
	return matched
}

// Function to check if the line contains a deadlock
func containsDeadlock(line string) bool {
	pattern := `WL-000394|BEA-000394`
	matched, _ := regexp.MatchString(pattern, line)
	return matched
}

// Function to check if the line contains an unchecked exception
func containsUncheckedException(line string) bool {
	pattern := `WL-000337|BEA-000337`
	matched, _ := regexp.MatchString(pattern, line)
	return matched
}

// Function to remove duplicate lines from the diagnostics log
func removeDuplicates(fileName string) {
	cmd := exec.Command("awk", "!seen[$0]++", fileName)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error removing duplicates:", err)
		return
	}

	err = os.WriteFile(fileName, output, 0644)
	if err != nil {
		fmt.Println("Error writing to diagnostics log:", err)
	}
}

