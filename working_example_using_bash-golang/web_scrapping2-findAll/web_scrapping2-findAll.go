package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL to scrape
	url := "http://en.wikipedia.org/wiki/Comparison_of_text_editors"

	// Fetch the page
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching page:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Error: Status code %d while fetching URL", resp.StatusCode)
	}

	// Parse the page
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// Find the first table with class 'wikitable'
	table := doc.Find("table.wikitable").First()

	// Open CSV file for writing
	csvFile, err := os.Create("editors.csv")
	if err != nil {
		log.Fatal("Error creating CSV file:", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Loop through table rows
	table.Find("tr").Each(func(i int, row *goquery.Selection) {
		var csvRow []string
		// Loop through all 'td' and 'th' cells in the row
		row.Find("td, th").Each(func(j int, cell *goquery.Selection) {
			// Get the cell text and append to the csvRow slice
			csvRow = append(csvRow, cell.Text())
		})
		// Write the row to the CSV
		err := writer.Write(csvRow)
		if err != nil {
			log.Fatal("Error writing to CSV:", err)
		}
	})

	fmt.Println("CSV file has been created successfully.")
}

