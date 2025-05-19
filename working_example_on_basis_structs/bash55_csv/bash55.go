package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/gocolly/colly"
)

// Middleware struct to represent our middleware
type Middleware struct {
	c *colly.Collector
}

// NewMiddleware initializes a new Middleware instance
func NewMiddleware() *Middleware {
	c := colly.NewCollector(
		colly.AllowedDomains("wikipedia.org"),
		colly.URLFilters(
			regexp.MustCompile(`^https://en.wikipedia.org/wiki/.*$`),
		),
	)

	m := &Middleware{c}

	// Setup middleware-like behavior
	c.OnRequest(func(r *colly.Request) {
		m.processStartRequest(r)
	})

	c.OnResponse(func(r *colly.Response) {
		m.processSpiderOutput(r)
	})

	c.OnError(func(r *colly.Response, err error) {
		m.processSpiderException(r, err)
	})

	return m
}

// processStartRequest simulates the process_start_requests method from Scrapy
func (m *Middleware) processStartRequest(r *colly.Request) {
	// This function can be used to modify requests before they are sent.
	fmt.Printf("Starting request: %s\n", r.URL)
}

// processSpiderOutput simulates the process_spider_output method from Scrapy
func (m *Middleware) processSpiderOutput(r *colly.Response) {
	// Process the response received from the server
	fmt.Printf("Received response for: %s\n", r.Request.URL)
	// You can further process the response as needed, e.g., extracting data
}

// processSpiderException simulates the process_spider_exception method from Scrapy
func (m *Middleware) processSpiderException(r *colly.Response, err error) {
	// Handle exceptions here
	log.Printf("Error occurred for request %s: %v\n", r.Request.URL, err)
}

// Example usage
func main() {
	middleware := NewMiddleware()

	// Define the main scraping behavior
	middleware.c.OnHTML("h1", func(e *colly.HTMLElement) {
		title := e.Text
		fmt.Printf("Title: %s\n", title)
	})

	// Start scraping from a specific URL
	startURL := "https://en.wikipedia.org/wiki/Benevolent_dictator_for_life"
	err := middleware.c.Visit(startURL)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for all scraping jobs to finish
	middleware.c.Wait()
}

