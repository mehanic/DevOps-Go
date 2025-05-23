package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// ReadDict reads a CSV file and returns a slice of maps (dictionaries)
func readDict(fileLocation string) ([]map[string]string, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allows variable number of fields per record

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, nil
	}

	headers := records[0]
	var result []map[string]string

	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, value := range record {
			row[headers[i]] = value
		}
		result = append(result, row)
	}

	return result, nil
}

// WriteDict writes a slice of maps to a CSV file
func writeDict(fileLocation string, dictionaries []map[string]string) error {
	file, err := os.Create(fileLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if len(dictionaries) == 0 {
		return nil
	}

	headers := make([]string, 0, len(dictionaries[0]))
	for key := range dictionaries[0] {
		headers = append(headers, key)
	}
	writer.Write(headers)

	for _, dictionary := range dictionaries {
		record := make([]string, len(headers))
		for i, header := range headers {
			record[i] = dictionary[header]
		}
		writer.Write(record)
	}

	return nil
}

// DictTest handles user input and processes CSV file operations
func dictTest() {
	var inputRows []map[string]string
	keepGoing := true

	for keepGoing {
		var name, lastName, age string
		fmt.Print("Name: ")
		fmt.Scan(&name)
		fmt.Print("Last Name: ")
		fmt.Scan(&lastName)
		fmt.Print("Age: ")
		fmt.Scan(&age)

		inputRows = append(inputRows, map[string]string{"name": name, "last_name": lastName, "age": age})

		var uiKeepGoing string
		fmt.Print("continue? (y/N): ")
		fmt.Scan(&uiKeepGoing)
		if strings.ToLower(uiKeepGoing) != "y" {
			keepGoing = false
		}
	}

	fmt.Println(inputRows)

	if err := writeDict("dict.csv", inputRows); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	writtenValue, err := readDict("dict.csv")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println(writtenValue)
}

func main() {
	dictTest()
}

