package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Data to write to CSV
	rows := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}

	// Write data to CSV file
	writeFile := "my.csv1"
	file, err := os.Create(writeFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing record to file:", err)
			return
		}
	}

	// Read data from CSV file
	readFile := "my.csv1"
	file, err = os.Open(readFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading record:", err)
			return
		}
		fmt.Println(record)
	}
}

