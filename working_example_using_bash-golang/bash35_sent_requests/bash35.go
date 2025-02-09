package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL to scrape
	url := "http://ualr.edu/safety/home/campus-safety-links/crime-statistics/"

	// Send HTTP request to the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch the page: %s", resp.Status)
	}

	// Load the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Open the output file
	file, err := os.Create("crime-stats-output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get the first table (Crime Statistics for UALR Main Campus)
	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		if i == 0 { // Only write the first table
			// Write the text of the table to the file
			tableText := table.Text()
			_, err := file.WriteString(tableText)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	fmt.Println("Crime statistics have been written to crime-stats-output.txt")
}

