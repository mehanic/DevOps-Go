package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// URLs to scrape
var urls = []string{
	"http://en.wikipedia.org/wiki/Python_%28programming_language%29",
	"https://en.wikipedia.org/wiki/Functional_programming",
	"https://en.wikipedia.org/wiki/Monty_Python",
}

// scrapeURL handles the HTTP request and HTML parsing for a single URL
func scrapeURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching URL %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != 200 {
		log.Printf("Received non-200 status code for URL %s: %d", url, resp.StatusCode)
		return
	}

	// Parse the HTML using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML for URL %s: %v", url, err)
		return
	}

	// Extract the title (similar to response.css('h1::text'))
	title := doc.Find("h1").Text()

	// Output the results
	fmt.Printf("URL is: %s\n", url)
	fmt.Printf("Title is: %s\n", title)
}

func main() {
	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Start scraping each URL concurrently
	for _, url := range urls {
		wg.Add(1)
		go scrapeURL(url, &wg)
	}

	// Wait for all scraping tasks to complete
	wg.Wait()
}

