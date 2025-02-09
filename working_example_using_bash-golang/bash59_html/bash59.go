package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	// Compile the regex pattern for image URLs
	imagePattern := regexp.MustCompile(`\.\./img/gifts/img.*\.jpg`)

	// On every HTML element with an <img> tag, extract the src attribute
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		src := e.Attr("src")
		if imagePattern.MatchString(src) {
			fmt.Println(src)
		}
	})

	// Start scraping the specified URL
	startURL := "http://www.pythonscraping.com/pages/page3.html"
	err := c.Visit(startURL)
	if err != nil {
		log.Fatal(err)
	}
}

