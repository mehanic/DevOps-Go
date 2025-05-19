package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	// Step 1: Create log file and write values
	createLogFile()

	// Step 2: Unique sorted values
	uniqueSortedValues()

	// Step 3: Count occurrences of each value
	countOccurrences()

	// Step 4: Show duplicate values
	duplicateValues()

	// Step 5: Show unique values
	uniqueValues()

	// Step 6: Character count
	characterCount("KSLDHGLKGHEIGUEFDDIFH")
}

func createLogFile() {
	if _, err := os.Stat("log"); err == nil {
		// Remove the log file if it exists
		err = os.Remove("log")
		if err != nil {
			fmt.Println("Error removing log file:", err)
		}
	}

	// Create a new log file and write numbers to it
	f, err := os.Create("log")
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return
	}
	defer f.Close()

	values := []string{"1", "2", "3", "4", "4"}
	for _, value := range values {
		f.WriteString(value + "\n")
	}
}

func uniqueSortedValues() {
	content, err := ioutil.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	uniqueValues := make(map[string]struct{})

	// Collect unique values
	for _, line := range lines {
		if line != "" {
			uniqueValues[line] = struct{}{}
		}
	}

	// Convert to slice and sort
	var sortedUnique []string
	for key := range uniqueValues {
		sortedUnique = append(sortedUnique, key)
	}
	sort.Strings(sortedUnique)

	fmt.Println("Unique Sorted Values:", sortedUnique)
}

func countOccurrences() {
	content, err := ioutil.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	valueCount := make(map[string]int)

	// Count occurrences
	for _, line := range lines {
		if line != "" {
			valueCount[line]++
		}
	}

	fmt.Println("Count of Occurrences:")
	for key, count := range valueCount {
		fmt.Printf("%d %s\n", count, key)
	}
}

func duplicateValues() {
	content, err := ioutil.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	valueCount := make(map[string]int)

	// Count occurrences
	for _, line := range lines {
		if line != "" {
			valueCount[line]++
		}
	}

	var duplicates []string
	for key, count := range valueCount {
		if count > 1 {
			duplicates = append(duplicates, key)
		}
	}

	fmt.Println("Duplicate Values:", duplicates)
}

func uniqueValues() {
	content, err := ioutil.ReadFile("log")
	if err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	valueCount := make(map[string]int)

	// Count occurrences
	for _, line := range lines {
		if line != "" {
			valueCount[line]++
		}
	}

	var uniques []string
	for key, count := range valueCount {
		if count == 1 {
			uniques = append(uniques, key)
		}
	}

	fmt.Println("Unique Values:", uniques)
}

func characterCount(input string) {
	characters := strings.Split(input, "")
	charCount := make(map[string]int)

	// Count occurrences of each character
	for _, char := range characters {
		if char != "" {
			charCount[char]++
		}
	}

	// Prepare output
	var output []string
	for char, count := range charCount {
		output = append(output, fmt.Sprintf("%d %s", count, char))
	}

	// Sort the output by character
	sort.Strings(output)

	// Join and print the output
	result := strings.Join(output, " ")
	fmt.Println("Character Count Output:", result)
}

