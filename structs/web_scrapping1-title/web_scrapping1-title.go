package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

// Article represents a scraped article
type Article struct {
	Title string
}

func main() {
	// Create a new collector
	c := colly.NewCollector(
		// Limit the domain to Wikipedia
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// Slice to store scraped articles
	var articles []Article

	// On every <h1> element found, extract the title
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		// Create a new Article item
		article := Article{
			Title: e.Text,
		}

		// Print the title
		fmt.Println("Title is:", article.Title)

		// Add the article to the list
		articles = append(articles, article)
	})

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping on the defined URLs
	startURLs := []string{
		"http://en.wikipedia.org/wiki/Main_Page",
		"http://en.wikipedia.org/wiki/Python_%28programming_language%29",
	}

	for _, url := range startURLs {
		err := c.Visit(url)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Scrape finished
	fmt.Println("Scraping finished")
}

