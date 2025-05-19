package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL to scrape
	url := "http://en.wikipedia.org/wiki/Python_(programming_language)"

	// Fetch the page
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching page:", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != 200 {
		log.Fatalf("Error: Status code %d while fetching URL", resp.StatusCode)
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error parsing HTML:", err)
	}

	// Extract content from the div with id "mw-content-text"
	content := doc.Find("div#mw-content-text").Text()

	// Print the extracted content
	fmt.Println(content)
}

