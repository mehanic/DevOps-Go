package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Opening workbook...")

	// Check if the file exists
	if _, err := os.Stat("censuspopdata.xlsx"); os.IsNotExist(err) {
		fmt.Println("Error: File does not exist.")
		return
	}

	// If the file exists, proceed with processing
	fmt.Println("Processing the Excel file...")
	// Further steps for processing the file would follow...
}

