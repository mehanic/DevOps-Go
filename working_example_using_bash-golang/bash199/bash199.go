package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"gopkg.in/yaml.v2"
)

func main() {
	// Open the Excel workbook
	fmt.Println("Opening workbook...")
	wb, err := excelize.OpenFile("censuspopdata.xlsx")
	if err != nil {
		log.Fatalf("Failed to open workbook: %v", err)
	}

	// Get the sheet by name
	sheetName := "Population by Census Tract"
	rows, err := wb.GetRows(sheetName)
	if err != nil {
		log.Fatalf("Failed to get rows: %v", err)
	}

	// Data structure to store county data
	countyData := make(map[string]map[string]map[string]int)

	// Start reading rows (skip header, starting from row 2)
	fmt.Println("Reading rows...")
	for i, row := range rows {
		if i == 0 { // Skip the header
			continue
		}

		state := row[1]
		county := row[2]
		pop, err := strconv.Atoi(row[3]) // Convert population to int
		if err != nil {
			log.Printf("Failed to convert population for row %d: %v", i+1, err)
			continue
		}

		// Ensure the state key exists
		if _, exists := countyData[state]; !exists {
			countyData[state] = make(map[string]map[string]int)
		}

		// Ensure the county key exists within the state
		if _, exists := countyData[state][county]; !exists {
			countyData[state][county] = map[string]int{"tracts": 0, "pop": 0}
		}

		// Increment the number of tracts and population
		countyData[state][county]["tracts"] += 1
		countyData[state][county]["pop"] += pop
	}

	// Write results to a file
	fmt.Println("Writing results...")
	resultFile, err := os.Create("census2010.yml")
	if err != nil {
		log.Fatalf("Failed to create result file: %v", err)
	}
	defer resultFile.Close()

	// Convert countyData to YAML and write to file
	encoder := yaml.NewEncoder(resultFile)
	err = encoder.Encode(countyData)
	if err != nil {
		log.Fatalf("Failed to write data to file: %v", err)
	}
	encoder.Close()

	fmt.Println("Done.")
}

