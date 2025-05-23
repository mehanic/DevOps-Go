package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Create or open a CSV file for writing
	csvFile, err := os.Create("test.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer csvFile.Close()

	// Initialize CSV writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush() // Ensure all buffered data is written to the file

	// Write the header row
	err = writer.Write([]string{"number", "number plus 2", "number times 2"})
	if err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}

	// Write data rows
	for i := 0; i < 10; i++ {
		row := []string{
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%d", i+2),
			fmt.Sprintf("%d", i*2),
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Println("Error writing to CSV:", err)
			return
		}
	}
}

