package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

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

	// Slice to store crawled articles
	var articles []Article

	// On every valid page that matches the filter, we extract data
	c.OnHTML("html", func(e *colly.HTMLElement) {
		// Check if the URL contains a colon (:) to exclude special pages
		if strings.Contains(e.Request.URL.String(), ":") {
			return
		}

		article := Article{
			URL:   e.Request.URL.String(),
			Title: e.ChildText("h1"),                     // Extract the title
			Text:  e.ChildText("#mw-content-text"),       // Extract article text
			LastUpdated: strings.TrimPrefix(
				e.ChildText("#footer-info-lastmod"),
				"This page was last edited on ",
			), // Extract last updated info
		}

		// Print out the collected article data
		fmt.Printf("Scraped URL: %s\nTitle: %s\nLast Updated: %s\n", article.URL, article.Title, article.LastUpdated)

		// Add the article to the list
		articles = append(articles, article)
	})

	// Follow links that match the pattern
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Only follow internal Wikipedia links that do not contain a colon
		if matched, _ := regexp.MatchString(`^/wiki/[^:]*$`, link); matched {
			// Visit the next link
			e.Request.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// Start scraping the first URL
	startURL := "https://en.wikipedia.org/wiki/Benevolent_dictator_for_life"
	err := c.Visit(startURL)
	if err != nil {
		log.Fatal(err)
	}

	// Wait until all requests are done
	c.Wait()
}