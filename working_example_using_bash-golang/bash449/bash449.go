package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

const (
	baseURL = "https://imdb.com"
	dataDir = "data"
)

// fetchHTML fetches the HTML content of a given URL.
func fetchHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

// extractLinks extracts links from the given HTML content.
func extractLinks(htmlContent []byte) ([]string, error) {
	doc, err := html.Parse(bytes.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links, nil
}

// fetchLinks recursively fetches links from the base URL.
func fetchLinks(url string, depth int, wg *sync.WaitGroup, linksChan chan<- string) {
	defer wg.Done()
	if depth == 0 {
		return
	}

	htmlContent, err := fetchHTML(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	links, err := extractLinks(htmlContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, link := range links {
		if strings.HasPrefix(link, "/") {
			link = baseURL + link
		}
		linksChan <- link
		fetchLinks(link, depth-1, wg, linksChan)
	}
}

// uniqueLinks collects unique links from the channel.
func uniqueLinks(linksChan <-chan string, doneChan <-chan struct{}) []string {
	linksSet := make(map[string]struct{})
	for {
		select {
		case link := <-linksChan:
			linksSet[link] = struct{}{}
		case <-doneChan:
			var links []string
			for link := range linksSet {
				links = append(links, link)
			}
			sort.Strings(links) // Sort the links
			return links
		}
	}
}

func main() {
	// Create the data directory
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		fmt.Printf("Error creating data directory: %v\n", err)
		return
	}

	// Initialize channels and wait group
	linksChan := make(chan string)
	doneChan := make(chan struct{})
	var wg sync.WaitGroup

	// Start fetching links
	wg.Add(1)
	go func() {
		defer close(linksChan)
		fetchLinks(baseURL, 5, &wg, linksChan)
	}()

	// Collect unique links
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	uniqueLinksList := uniqueLinks(linksChan, doneChan)

	// Write links to CSV file
	file, err := os.Create("links_final.csv")
	if err != nil {
		fmt.Printf("Error creating CSV file: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, link := range uniqueLinksList {
		if strings.HasPrefix(link, "http") {
			if err := writer.Write([]string{link}); err != nil {
				fmt.Printf("Error writing to CSV: %v\n", err)
			}
		}
	}

	// Clean up
	if err := os.RemoveAll(dataDir); err != nil {
		fmt.Printf("Error removing data directory: %v\n", err)
		return
	}

	fmt.Println("Links extracted and saved to links_final.csv")
}

