package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	// Load the workbook
	wb, err := xlsx.OpenFile("produceSales.xlsx")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	// Get the first sheet
	sheet := wb.Sheets[0]

	// The produce types and their updated prices
	priceUpdates := map[string]float64{
		"Garlic": 3.07,
		"Celery": 1.19,
		"Lemon":  1.27,
	}

	// Loop through the rows and update the prices.
	for rowNum := 1; rowNum < sheet.MaxRow; rowNum++ { // start from 1 to skip the header row
		produceName := sheet.Cell(rowNum, 0).String() // first column
		if newPrice, exists := priceUpdates[produceName]; exists {
			// Update the price in the second column
			sheet.Cell(rowNum, 1).SetFloat(newPrice) // second column
		}
	}

	// Save the updated workbook
	err = wb.Save("updatedProduceSales.xlsx")
	if err != nil {
		log.Fatalf("Error saving file: %v", err)
	}

	fmt.Println("Updated prices saved to 'updatedProduceSales.xlsx'")
}

