package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

// Article struct to hold the scraped data
type Article struct {
	URL         string
	Title       string
	Text        string
	LastUpdated string
}

func main() {
	// Create a new collector with allowed domains set to wikipedia.org
	c := colly.NewCollector(
		colly.AllowedDomains("wikipedia.org"),
	)

	// Slice to store the articles
	var articles []Article
	// Mutex to handle concurrent access to articles slice
	var mutex sync.Mutex

	// On every valid page that matches the URL filter, extract the article data
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// Check if the URL contains a colon (:) to exclude special pages
		if strings.Contains(e.Request.URL.String(), ":") {
			return
		}

		mutex.Lock() // Locking for concurrent writes
		defer mutex.Unlock()

		// Extract the title
		title := e.ChildText("h1")

		// Extract the main content by iterating over <p> tags within #mw-content-text
		var textBuilder strings.Builder
		e.ForEach("#mw-content-text p", func(_ int, el *colly.HTMLElement) {
			textBuilder.WriteString(el.Text + " ")
		})
		text := strings.TrimSpace(textBuilder.String())

		// Extract the last updated date
		lastUpdated := e.ChildText("#footer-info-lastmod")
		lastUpdated = strings.TrimPrefix(lastUpdated, "This page was last edited on ")

		// Create the article
		article := Article{
			URL:         e.Request.URL.String(),
			Title:       title,
			Text:        text,
			LastUpdated: lastUpdated,
		}

		// Add article to the list
		articles = append(articles, article)

		// Print the extracted data
		fmt.Printf("Scraped URL: %s\nTitle: %s\nLast Updated: %s\n", article.URL, article.Title, article.LastUpdated)
	})

	// Follow all valid Wikipedia links that match the /wiki/ pattern
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Only follow internal Wikipedia links that do not contain a colon
		if matched, _ := regexp.MatchString(`^/wiki/[^:]*$`, link); matched {
			e.Request.Visit(e.Request.AbsoluteURL(link)) // Visit the link
		}
	})

	// Log any errors encountered during scraping
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with error:", err)
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