package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Send a GET request to the URL
	response, err := http.Get("http://www.pythonscraping.com/pages/page3.html")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Check if the response status is OK
	if response.StatusCode != 200 {
		log.Fatalf("Error: status code %d", response.StatusCode)
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the table with the specified id
	table := doc.Find("table#giftList")

	// Get the first table row (tr)
	firstRow := table.Find("tr").First()

	// Iterate through the next siblings of the first row
	firstRow.Each(func(i int, s *goquery.Selection) {
		for sibling := s.Nodes[0].NextSibling; sibling != nil; sibling = sibling.NextSibling {
			// Convert the sibling node to a goquery selection
			siblingSelection := goquery.NewDocumentFromNode(sibling)

			// Print the HTML of the sibling (or you can print its text if needed)
			html, err := goquery.OuterHtml(siblingSelection.Selection)
			if err == nil {
				fmt.Println(html)
			}
		}
	})
}

