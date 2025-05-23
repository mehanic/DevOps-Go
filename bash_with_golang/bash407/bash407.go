package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Function to read a file and return its content as a string
func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Function to write a string to a file
func writeFile(filename, content string) error {
	return ioutil.WriteFile(filename, []byte(content), 0644)
}

// Function to process CSV content and return as a formatted string
func processCSV(inputFile, delimiter string) (string, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = rune(delimiter[0]) // Set the delimiter

	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	for _, record := range records {
		// Join the record with delimiter and add to the builder
		sb.WriteString(strings.Join(record, delimiter))
		sb.WriteString("\n") // Add a newline after each record
	}
	return sb.String(), nil
}

func main() {
	// File paths
	outputFile := "words.xml"
	inputFile := "words.csv"
	
	// Read template files
	xmlHdr, err := readFile("xml-parent.tpl")
	if err != nil {
		log.Fatalf("Error reading xml-parent.tpl: %v", err)
	}
	srtContr, err := readFile("word.tpl")
	if err != nil {
		log.Fatalf("Error reading word.tpl: %v", err)
	}
	endContr := "</root>" // Replace this with the appropriate trailer

	// Setup header
	outputContent := xmlHdr + srtContr

	// Process CSV and append content
	csvContent, err := processCSV(inputFile, ",")
	if err != nil {
		log.Fatalf("Error processing CSV file: %v", err)
	}
	outputContent += csvContent

	// Append trailer
	outputContent += endContr

	// Write the final content to the output file
	if err := writeFile(outputFile, outputContent); err != nil {
		log.Fatalf("Error writing to output file: %v", err)
	}

	// Output the result
	fmt.Println("Output content:")
	fmt.Println(outputContent)
}

