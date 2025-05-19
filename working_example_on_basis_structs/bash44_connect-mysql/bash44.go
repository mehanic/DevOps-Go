package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

// Article structure
type Article struct {
	Title string
	Path  string
}

// Channel for storing articles
var queue = make(chan Article, 100)

// Mutex for handling visited URLs concurrency
var visited = struct {
	m map[string]bool
	sync.Mutex
}{m: make(map[string]bool)}

// MySQL connection setup
func connectDB() (*sql.DB, error) {
	dsn := "root:@unix(/tmp/mysql.sock)/wiki_threads?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Function to store articles in MySQL
func storage() {
	db, err := connectDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	for article := range queue {
		// Store article in DB
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM pages WHERE path = ?", article.Path).Scan(&count)
		if err != nil {
			fmt.Println("Error querying database:", err)
			continue
		}

		if count == 0 {
			fmt.Printf("Storing article: %s\n", article.Title)
			_, err := db.Exec("INSERT INTO pages (title, path) VALUES (?, ?)", article.Title, article.Path)
			if err != nil {
				fmt.Println("Error inserting article:", err)
			}
		} else {
			fmt.Printf("Article already exists: %s\n", article.Title)
		}
	}
}

// Fetch links from the article page
func getLinks(threadName string, doc *goquery.Document) []string {
	fmt.Printf("Getting links in %s\n", threadName)
	var links []string
	// Regular expression to match Wikipedia article URLs
	doc.Find("div#bodyContent a[href^='/wiki/']").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			visited.Lock()
			if !visited.m[href] {
				links = append(links, href)
			}
			visited.Unlock()
		}
	})
	return links
}

// Function to scrape a Wikipedia article
func scrapeArticle(threadName, path string) {
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

	// Extract article title
	title := doc.Find("h1").Text()
	fmt.Printf("Added %s for storage in thread %s\n", title, threadName)

	// Put the article into the storage queue
	queue <- Article{Title: title, Path: path}

	// Get the links from the current page
	links := getLinks(threadName, doc)

	// If there are new links, randomly pick one and scrape it
	if len(links) > 0 {
		newArticle := links[rand.Intn(len(links))]
		scrapeArticle(threadName, newArticle)
	}
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Start goroutines for scraping
	go scrapeArticle("Thread 1", "/wiki/Kevin_Bacon")
	go scrapeArticle("Thread 2", "/wiki/Monty_Python")

	// Start goroutine for storage
	go storage()

	// Run indefinitely
	select {}
}

