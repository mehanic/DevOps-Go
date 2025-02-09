package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL to scrape
	url := "http://thomaswallace.net/courses/internet-technologies/fall-2015/"

	// Send HTTP request to the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Load the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find all article elements with class "post"
	doc.Find("article.post").Each(func(i int, post *goquery.Selection) {
		// Iterate over all children of each post
		post.Children().Each(func(i int, child *goquery.Selection) {
			// Print each child element
			fmt.Println(child.Text())
		})
		// Print a newline after each post
		fmt.Println("\n")
	})
}