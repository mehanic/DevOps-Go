package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"io"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL to scrape
	url := "http://www.thomaswallace.net"

	// Send HTTP request to the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch the page: %s", resp.Status)
	}

	// Read the content of the response
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the raw HTML content
	fmt.Println(string(content))

	// Load the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	if err != nil {
		log.Fatal(err)
	}

	// Print the prettified HTML (formatted version)
	fmt.Println(doc.Find("html").Html())

	// Get and print the title tag
	title := doc.Find("title")
	fmt.Println(title)

	// Get and print the title string
	titleString := title.Text()
	fmt.Println(titleString)

	// Get and print the first paragraph (<p>) tag
	firstParagraph := doc.Find("p").First()
	fmt.Println(firstParagraph)

	// Get and print the first anchor (<a>) tag
	firstAnchor := doc.Find("a").First()
	fmt.Println(firstAnchor)
}

