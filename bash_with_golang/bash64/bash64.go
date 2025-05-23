package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Send a GET request to the URL
	response, err := http.Get("http://www.pythonscraping.com/pages/page2.html")
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

	// Find all tags with exactly two attributes
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		// Get the attributes of the tag
		attributes := s.Nodes[0].Attr
		if len(attributes) == 2 {
			// Print the HTML of the tag
			html, err := goquery.OuterHtml(s)
			if err != nil {
				log.Println("Error getting tag HTML:", err)
				return
			}
			fmt.Println(html)
		}
	})
}

