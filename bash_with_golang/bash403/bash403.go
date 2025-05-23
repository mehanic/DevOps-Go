package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Function to read CSV data and return as a slice of strings
func readCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	return reader.ReadAll()
}

// Function to write CSV data from a slice of strings
func writeCSV(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	return writer.WriteAll(data) // Write all data at once
}

func main() {
	gbCSV := "testdata/garbage.csv"
	emCSV := "testdata/employees.csv"

	// Process garbage.csv
	garbageData, err := readCSV(gbCSV)
	if err != nil {
		log.Fatalf("Error reading %s: %v", gbCSV, err)
	}

	fmt.Printf("We have %d rows in %s\n", len(garbageData), gbCSV)

	// Strip garbage: remove spaces and make all upper case
	for i, row := range garbageData {
		for j := range row {
			row[j] = strings.ReplaceAll(row[j], " ", "") // Remove spaces
			row[j] = strings.ToUpper(row[j])             // Convert to uppercase
		}
		garbageData[i] = row
		fmt.Println(garbageData[i])
	}

	// Process employees.csv
	employeeData, err := readCSV(emCSV)
	if err != nil {
		log.Fatalf("Error reading %s: %v", emCSV, err)
	}

	// Add a '#' at the beginning of each line
	for i, row := range employeeData {
		for j := range row {
			row[j] = "#" + row[j] // Prepend '#' to each field
		}
		employeeData[i] = row
		fmt.Println(employeeData[i])
	}

	// Modify the employees.csv directly by replacing Bob with Robert
	modifyEmployeeFile(emCSV)

	// Remove the birthdate field (5th column, index 4)
	removeBirthdayField(emCSV)
}

// Function to modify the employee CSV by replacing Bob with Robert
func modifyEmployeeFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading %s: %v", filename, err)
	}

	// Replace "Bob" with "Robert"
	modifiedData := strings.ReplaceAll(string(data), "Bob", "Robert")

	// Prepend '#' to each line and create a new string
	lines := strings.Split(modifiedData, "\n")
	for i := range lines {
		lines[i] = "#" + lines[i] // Prepend '#' to each line
	}

	// Write the modified lines back to the file
	err = ioutil.WriteFile(filename, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error writing to %s: %v", filename, err)
	}

	// Print the final content of the file
	fmt.Println("Final content of the modified employee CSV:")
	fmt.Println(strings.Join(lines, "\n")) // Join the lines back into a single string for printing
}

// Function to remove the birthday field from the employees CSV
func removeBirthdayField(filename string) {
	data, err := readCSV(filename)
	if err != nil {
		log.Fatalf("Error reading %s: %v", filename, err)
	}

	for i := range data {
		if len(data[i]) > 4 { // Ensure there's at least 5 fields
			data[i] = append(data[i][:4], data[i][5:]...) // Remove the 5th column
		}
	}

	err = writeCSV(filename, data)
	if err != nil {
		log.Fatalf("Error writing to %s: %v", filename, err)
	}
}

