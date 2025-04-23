package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Data to be written to CSV
	dictionaries := []map[string]string{
		{"age": "30", "name": "John", "last_name": "Doe"},
		{"age": "30", "name": "Jane", "last_name": "Doe"},
	}

	// Open CSV file for writing
	csvFile, err := os.Create("my.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write header
	headers := []string{"age", "name", "last_name"}
	if err := writer.Write(headers); err != nil {
		fmt.Println("Error writing header:", err)
		return
	}

	// Write records
	for _, record := range dictionaries {
		row := []string{
			record["age"],
			record["name"],
			record["last_name"],
		}
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}

	// Open CSV file for reading
	csvFile, err = os.Open("my.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
		return
	}

	// Print all records
	for _, record := range records {
		fmt.Println(record)
	}
}

