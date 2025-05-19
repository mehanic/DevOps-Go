package main

import (
	"fmt"
	"log"
	"net/http"
	

	"github.com/PuerkitoBio/goquery"
)

var pages = make(map[string]bool) // Set to track visited pages

// getLinks fetches links from the specified Wikipedia page
func getLinks(pageUrl string) {
	// Build the full URL
	fullUrl := "http://en.wikipedia.org" + pageUrl
	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Println("Failed to fetch URL:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: status code %d for URL: %s\n", resp.StatusCode, fullUrl)
		return
	}

	bs, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error loading HTML:", err)
		return
	}

	// Get the page title
	if title := bs.Find("h1").Text(); title != "" {
		fmt.Println("Page title:", title)
	} else {
		fmt.Println("No title found")
	}

	// Get the first paragraph
	firstParagraph := bs.Find("#mw-content-text p").First().Text()
	if firstParagraph != "" {
		fmt.Println("First paragraph:", firstParagraph)
	} else {
		fmt.Println("No first paragraph found")
	}

	// Get the edit link
	editLink, exists := bs.Find("#ca-edit span a").Attr("href")
	if exists {
		fmt.Println("Edit link:", editLink)
	} else {
		fmt.Println("No edit link found")
	}

	// Find and follow internal links
	bs.Find("a[href^=/wiki/]").Each(func(index int, element *goquery.Selection) {
		link, exists := element.Attr("href")
		if exists && !pages[link] {
			// We have encountered a new page
			pages[link] = true
			fmt.Println("----------------\n" + link)
			getLinks(link) // Recursive call
		}
	})
}

func main() {
	getLinks("") // Start from the main page
}

