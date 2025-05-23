package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "sync"
    "time"

    "github.com/PuerkitoBio/goquery"
)

// This will keep track of visited URLs
var visited = struct {
    m map[string]bool
    sync.Mutex
}{m: make(map[string]bool)}

// Function to fetch and parse links from a Wikipedia article
func getLinks(threadName string, doc *goquery.Document) []string {
    fmt.Printf("Getting links in %s\n", threadName)
    var links []string
    // Find all valid /wiki/ links
    doc.Find("div#bodyContent a[href^='/wiki/']").Each(func(i int, s *goquery.Selection) {
        href, exists := s.Attr("href")
        if exists && !visited.m[href] {
            links = append(links, href)
        }
    })
    return links
}

// Function to scrape a Wikipedia article
func scrapeArticle(threadName, path string) {
    // Lock the visited map to avoid race conditions
    visited.Lock()
    visited.m[path] = true
    visited.Unlock()

    resp, err := http.Get("http://en.wikipedia.org" + path)
    if err != nil {
        fmt.Printf("Error fetching %s: %v\n", path, err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Printf("Error: Status code %d\n", resp.StatusCode)
        return
    }

    // Parse the HTML using goquery
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        fmt.Printf("Error parsing document: %v\n", err)
        return
    }

    title := doc.Find("h1").Text()
    fmt.Printf("Scraping %s in thread %s\n", title, threadName)

    // Get the links on the current page
    links := getLinks(threadName, doc)

    // If there are new links, pick a random one and scrape it
    if len(links) > 0 {
        newArticle := links[rand.Intn(len(links))]
        fmt.Printf("New article: %s\n", newArticle)
        time.Sleep(5 * time.Second) // Simulate delay
        scrapeArticle(threadName, newArticle) // Recursively scrape the next article
    }
}

func main() {
    // Initialize random seed
    rand.Seed(time.Now().UnixNano())

    // Start goroutines for scraping
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        defer wg.Done()
        scrapeArticle("Thread 1", "/wiki/Kevin_Bacon")
    }()

    go func() {
        defer wg.Done()
        scrapeArticle("Thread 2", "/wiki/Monty_Python")
    }()

    // Wait for the goroutines to finish
    wg.Wait()
}

