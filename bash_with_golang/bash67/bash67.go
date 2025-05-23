package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// getLinks retrieves all valid Wikipedia links from the specified article URL
func getLinks(articleUrl string) ([]string, error) {
	// Send a GET request to the Wikipedia article
	response, err := http.Get("http://en.wikipedia.org" + articleUrl)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check if the response status is OK
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch URL: status code %d", response.StatusCode)
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	// Find all links in the body content
	var links []string
	doc.Find("div#bodyContent a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists && regexp.MustCompile(`^(/wiki/)((?!:).)*$`).MatchString(href) {
			links = append(links, href)
		}
	})

	return links, nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Start with the first article
	links, err := getLinks("/wiki/Kevin_Bacon")
	if err != nil {
		log.Fatal(err)
	}

	// Continue retrieving links while there are still links available
	for len(links) > 0 {
		newArticle := links[rand.Intn(len(links))] // Get a random link
		fmt.Println(newArticle)
		links, err = getLinks(newArticle) // Retrieve links from the new article
		if err != nil {
			log.Fatal(err)
		}
	}
}

