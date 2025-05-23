package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var visited = struct {
	mu sync.Mutex
	m  map[string]bool
}{m: make(map[string]bool)}

// getLinks parses the HTML and returns valid Wikipedia article links
func getLinks(threadName string, doc *goquery.Document) []string {
	fmt.Printf("Getting links in %s\n", threadName)

	links := []string{}
	visited.mu.Lock()
	defer visited.mu.Unlock()

	doc.Find("div#bodyContent a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			match, _ := regexp.MatchString("^/wiki/((?!:).)*$", href)
			if match && !visited.m[href] {
				links = append(links, href)
			}
		}
	})

	return links
}

// scrapeArticle recursively scrapes articles, simulating the Python threading example
func scrapeArticle(wg *sync.WaitGroup, threadName, path string) {
	defer wg.Done()

	visited.mu.Lock()
	visited.m[path] = true
	visited.mu.Unlock()

	url := fmt.Sprintf("https://en.wikipedia.org%s", path)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Received non-200 response code for %s", url)
		return
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML for %s: %v", url, err)
		return
	}

	title := doc.Find("h1").Text()
	fmt.Printf("Scraping %s in %s\n", title, threadName)

	links := getLinks(threadName, doc)
	if len(links) > 0 {
		newArticle := links[rand.Intn(len(links))]
		fmt.Println("Next article:", newArticle)
		// Recursive call to scrape the next article
		time.Sleep(5 * time.Second)
		scrapeArticle(wg, threadName, newArticle)
	}
}

func main() {
	var wg sync.WaitGroup

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Start two goroutines to simulate threading
	wg.Add(2)
	go scrapeArticle(&wg, "Thread 1", "/wiki/Kevin_Bacon")
	go scrapeArticle(&wg, "Thread 2", "/wiki/Monty_Python")

	// Wait for all goroutines to finish
	wg.Wait()
}

