package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly")

// Article struct to hold the scraped article data
type Article struct {
	URL         string
	Title       string
	Text        string
	LastUpdated string
}

func main() {
	// Create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("wikipedia.org"),
		colly.URLFilters(
			regexp.MustCompile(`^https://en.wikipedia.org/wiki/.*$`), // Filter to allow only Wikipedia article pages
		),
	)

	// Mutex for concurrent writes to the articles slice
	var mutex sync.Mutex
	var articles []Article

	// On every valid page that matches the URL filter, extract the article data
	c.OnHTML("html", func(e *colly.HTMLElement) {
		mutex.Lock()
		defer mutex.Unlock()

		article := Article{
			URL:         e.Request.URL.String(),
			Title:       e.ChildText("h1"),
			Text:        strings.Join(e.ChildTexts("#mw-content-text p"), " "), // Join paragraphs into a single string
			LastUpdated: strings.TrimPrefix(e.ChildText("#footer-info-lastmod"), "This page was last edited on "), // Extract last updated date
		}

		// Add the article to the list
		articles = append(articles, article)

		// Print the extracted data
		fmt.Printf("URL: %s\nTitle: %s\nText: %s\nLast Updated: %s\n", article.URL, article.Title, article.Text, article.LastUpdated)
	})

	// Follow all links from the current page
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if matched, _ := regexp.MatchString(`^/wiki/((?!:).)*$`, link); matched {
			e.Request.Visit(e.Request.AbsoluteURL(link)) // Visit the link
		}
	})

	// Log any errors encountered during scraping
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL: %s failed with error: %v\n", r.Request.URL, err)
	})

	// Start scraping from the initial URL
	startURL := "https://en.wikipedia.org/wiki/Benevolent_dictator_for_life"
	err := c.Visit(startURL)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for all scraping jobs to finish
	c.Wait()

	// Output the total number of articles scraped
	fmt.Printf("Total articles scraped: %d\n", len(articles))
}

