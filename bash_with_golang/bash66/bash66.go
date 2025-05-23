package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Send a GET request to the URL
	response, err := http.Get("http://www.pythonscraping.com/pages/warandpeace.html")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Check if the response status is OK
	if response.StatusCode != 200 {
		log.Fatalf("Error: status code %d", response.StatusCode)
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find all <span> elements with the class "green"
	doc.Find("span.green").Each(func(index int, element *goquery.Selection) {
		// Get the text from each <span>
		name := element.Text()
		fmt.Println(name)
	})
}

