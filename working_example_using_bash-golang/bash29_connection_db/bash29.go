package main

import (
	"database/sql"
//	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

// Struct to hold database connection
type Database struct {
	Conn *sql.DB
}

// Function to connect to the database
func connectDB() (*Database, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/scraping" // Change the DSN according to your setup
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return &Database{Conn: db}, nil
}

// Function to store title and content in the database
func (db *Database) store(title string, content string) error {
	_, err := db.Conn.Exec("INSERT INTO pages (title, content) VALUES (?, ?)", title, content)
	return err
}

// Function to get all links from a Wikipedia page
func getLinks(articleUrl string, db *Database) ([]string, error) {
	resp, err := http.Get("http://en.wikipedia.org" + articleUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Get title
	title := doc.Find("h1").Find("span").Text()
	// Get content
	content := doc.Find("div#mw-content-text p").First().Text()
	err = db.store(title, content) // Store title and content
	if err != nil {
		return nil, err
	}

	// Find all links
	var links []string
	doc.Find("div#bodyContent a").Each(func(index int, item *goquery.Selection) {
		href, exists := item.Attr("href")
		if exists {
			match, _ := regexp.MatchString("^(/wiki/)((?!:).)*$", href)
			if match {
				links = append(links, href)
			}
		}
	})
	return links, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Conn.Close()

	links, err := getLinks("/wiki/Kevin_Bacon", db)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through links
	for len(links) > 0 {
		newArticle := links[rand.Intn(len(links))] // Randomly select a link
		fmt.Println(newArticle)
		links, err = getLinks(newArticle, db) // Get new links from the selected article
		if err != nil {
			log.Fatal(err)
		}
	}
}

