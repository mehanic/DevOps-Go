package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"	
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// On every HTML element with the id "giftList", extract its children
	c.OnHTML("#giftList", func(e *colly.HTMLElement) {
		// Iterate through the children of the table with id "giftList"
		e.DOM.Children().Each(func(i int, child *goquery.Selection) {
			// Print the HTML of each child element
			html, err := child.Html()
			if err != nil {
				log.Println("Error getting child HTML:", err)
				return
			}
			fmt.Println(html)
		})
	})

	// Start scraping the specified URL
	startURL := "http://www.pythonscraping.com/pages/page3.html"
	err := c.Visit(startURL)
	if err != nil {
		log.Fatal(err)
	}
}

