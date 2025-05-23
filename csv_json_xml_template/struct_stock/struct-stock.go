package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Struct to store row data similar to namedtuple in Python
type Stock struct {
	Symbol string
	Price  float64
	Date   string
	Time   string
	Change float64
	Volume int
}

// Function to read and write CSV files
func rwCSV() {
	// Read from CSV
	file, err := os.Open("stocks.csv")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read() // Read the headers
	if err != nil {
		log.Fatalf("Failed to read headers: %v", err)
	}
	fmt.Println("Headers:", headers)

	// Process rows
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read rows: %v", err)
	}
	for _, record := range records {
		fmt.Println("Row:", record)
	}

	// If you want to map each row to a struct, you can do that manually:
	file.Seek(0, 0) // Reset file pointer to start
	reader = csv.NewReader(file)
	_, _ = reader.Read() // Skip headers

	for {
		row, err := reader.Read()
		if err != nil {
			break // End of file reached
		}
		// Convert fields and assign them to the Stock struct
		price, _ := strconv.ParseFloat(row[1], 64)
		change, _ := strconv.ParseFloat(row[4], 64)
		volume, _ := strconv.Atoi(row[5])

		stock := Stock{
			Symbol: row[0],
			Price:  price,
			Date:   row[2],
			Time:   row[3],
			Change: change,
			Volume: volume,
		}
		fmt.Println("Stock Change:", stock.Change)
	}

	// Write to CSV
	headers = []string{"Symbol", "Price", "Date", "Time", "Change", "Volume"}
	rows := [][]string{
		{"AA", "39.48", "6/11/2007", "9:36am", "-0.18", "181800"},
		{"AIG", "71.38", "6/11/2007", "9:36am", "-0.15", "195500"},
		{"AXP", "62.58", "6/11/2007", "9:36am", "-0.46", "935000"},
	}

	outFile, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	err = writer.Write(headers)
	if err != nil {
		log.Fatalf("Failed to write headers: %v", err)
	}

	err = writer.WriteAll(rows)
	if err != nil {
		log.Fatalf("Failed to write rows: %v", err)
	}

	writer.Flush()
	fmt.Println("CSV file written successfully.")
}

func main() {
	rwCSV()
}

