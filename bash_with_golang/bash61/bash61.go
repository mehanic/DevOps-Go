package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// On every HTML element with the id "text", extract the text
	c.OnHTML("#text", func(e *colly.HTMLElement) {
		// Print the text content of the element
		fmt.Println(e.Text)
	})

	// Start scraping the specified URL
	startURL := "http://www.pythonscraping.com/pages/warandpeace.html"
	err := c.Visit(startURL)
	if err != nil {
		log.Fatal(err)
	}
}

