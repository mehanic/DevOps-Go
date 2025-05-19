package main

import (
//	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

// Function to clean garbage CSV
func cleanGarbageCSV(filePath string) [][]string {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s, error: %v", filePath, err)
	}
	defer file.Close()

	// Read the CSV content
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV: %v", err)
	}

	fmt.Printf("We have %d rows in %s\n", len(records), filePath)

	// Process each row
	for i, record := range records {
		// Join the record into a single string
		line := strings.Join(record, ",")
		// Remove spaces
		line = strings.ReplaceAll(line, " ", "")
		// Remove the last character
		if len(line) > 0 {
			line = line[:len(line)-1]
		}
		// Convert to uppercase
		records[i] = []string{strings.ToUpper(line)}
		fmt.Println(records[i][0])
	}
	return records
}

// Function to update employees CSV
func updateEmployeesCSV(filePath string) [][]string {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s, error: %v", filePath, err)
	}
	defer file.Close()

	// Read the CSV content
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV: %v", err)
	}

	fmt.Printf("\nWe have %d rows in %s\n", len(records), filePath)

	// Prepend '#' to the first field of each record
	for i := range records {
		records[i][0] = "#" + records[i][0]
		fmt.Println(records[i][0])
	}

	// Change "Bob" to "Robert"
	fmt.Println("\nLet's make Bob, Robert!")
	for i := range records {
		records[i] = updateName(records[i], "Bob", "Robert")
		fmt.Println(strings.Join(records[i], ","))
	}

	// Remove the birthday column (index 4)
	fmt.Println("\nLet's remove the column: birthday (1-31)")
	for i := range records {
		records[i] = removeColumn(records[i], 4)
		fmt.Println(strings.Join(records[i], ","))
	}

	return records
}

// Update name from oldName to newName
func updateName(record []string, oldName, newName string) []string {
	for i, field := range record {
		if field == oldName {
			record[i] = newName
		}
	}
	return record
}

// Remove a column from the record
func removeColumn(record []string, columnToRemove int) []string {
	if columnToRemove < 0 || columnToRemove >= len(record) {
		return record // Return the record unchanged if the index is out of bounds
	}

	// Create a new slice to hold the updated record
	var updatedRecord []string
	for i, field := range record {
		if i != columnToRemove {
			updatedRecord = append(updatedRecord, field)
		}
	}
	return updatedRecord
}

func main() {
	gbCSV := "testdata/garbage.csv"
	emCSV := "testdata/employees.csv"

	// Clean garbage CSV
	cleanGarbageCSV(gbCSV)

	// Update employees CSV
	updateEmployeesCSV(emCSV)
}

