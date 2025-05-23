package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	reqFile := "new_info.csv"

	// Open the file
	fo, err := os.Open(reqFile)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer fo.Close()

	// Create a new CSV reader with '|' as the delimiter
	reader := csv.NewReader(fo)
	reader.Comma = '|'

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %s", err)
	}

	// Print each record
	for _, each := range records {
		fmt.Println(each)
	}
}
