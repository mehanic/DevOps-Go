package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Fetch the webpage
	response, err := http.Get("http://www.pythonscraping.com/pages/page1.html")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode != 200 {
		log.Fatalf("Error: Status code %d", response.StatusCode)
	}

	// Parse the HTML using goquery
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print <h1> using different approaches (similar to the BeautifulSoup script)
	fmt.Println("First <h1>:")
	doc.Find("h1").Each(func(index int, item *goquery.Selection) {
		fmt.Println(item.Text())
	})

	fmt.Println("\n<h1> inside <html><body>:")
	doc.Find("html body h1").Each(func(index int, item *goquery.Selection) {
		fmt.Println(item.Text())
	})
}

