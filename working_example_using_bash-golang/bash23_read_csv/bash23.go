package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	gbCSV := "testdata/garbage.csv"
	emCSV := "testdata/employees.csv"

	// Process garbage.csv
	garbageRecords := readCSV(gbCSV)
	fmt.Printf("We have %d rows in %s\n", len(garbageRecords), gbCSV)

	// Strip garbage - remove spaces and process records
	for i, record := range garbageRecords {
		// Remove spaces
		garbageRecords[i] = strings.ReplaceAll(record, " ", "")
		// Remove the last character
		if len(garbageRecords[i]) > 0 {
			garbageRecords[i] = garbageRecords[i][:len(garbageRecords[i])-1]
		}
		// Convert to uppercase
		garbageRecords[i] = strings.ToUpper(garbageRecords[i])
		fmt.Println(garbageRecords[i])
	}

	// Process employees.csv
	employeeRecords := readCSV(emCSV)
	fmt.Printf("We have %d rows in %s\n", len(employeeRecords), emCSV)

	// Prepend '#' to each employee record
	for i := range employeeRecords {
		employeeRecords[i] = "#" + employeeRecords[i]
		fmt.Println(employeeRecords[i])
	}

	// Change "Bob" to "Robert"
	fmt.Println("\nLet's make Bob, Robert!")
	for i := range employeeRecords {
		employeeRecords[i] = strings.ReplaceAll(employeeRecords[i], "Bob", "Robert")
		fmt.Println(employeeRecords[i])
	}

	// Remove the day of birth column (5th column, which is index 4)
	fmt.Println("\nLet's remove the column: birthday (1-31)")
	const columnToRemove = 4
	for _, record := range employeeRecords {
		fields := strings.Split(record, ",")
		var newRecord []string
		for j, field := range fields {
			if j != columnToRemove {
				newRecord = append(newRecord, field)
			}
		}
		fmt.Println(strings.Join(newRecord, ","))
	}
}

// readCSV reads a CSV file and returns its contents as a slice of strings.
func readCSV(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return nil
	}
	defer file.Close()

	var records []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		records = append(records, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
	}
	return records
}

