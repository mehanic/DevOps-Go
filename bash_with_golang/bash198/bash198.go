package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Create the 'headerRemoved' directory if it doesn't exist
	err := os.MkdirAll("headerRemoved", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Loop through every file in the current working directory
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		// Process only CSV files
		if filepath.Ext(file.Name()) != ".csv" {
			continue
		}

		fmt.Println("Removing header from", file.Name(), "...")

		// Open the CSV file
		inFile, err := os.Open(file.Name())
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer inFile.Close()

		reader := csv.NewReader(inFile)

		// Skip the header (first row)
		var rows [][]string
		skipHeader := true
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error reading CSV:", err)
				break
			}
			if skipHeader {
				skipHeader = false
				continue
			}
			rows = append(rows, record)
		}

		// Create the output CSV file in the 'headerRemoved' folder
		outFilePath := filepath.Join("headerRemoved", file.Name())
		outFile, err := os.Create(outFilePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			continue
		}
		defer outFile.Close()

		writer := csv.NewWriter(outFile)
		defer writer.Flush()

		// Write the data (without header) to the new CSV file
		for _, row := range rows {
			err := writer.Write(row)
			if err != nil {
				fmt.Println("Error writing to CSV:", err)
				break
			}
		}
	}
}

