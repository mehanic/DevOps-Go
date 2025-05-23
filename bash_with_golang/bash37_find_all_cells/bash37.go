package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL to scrape
	url := "https://www.congress.gov/resources/display/content/Most-Viewed+Bills"

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

	// Load the HTML document using goquery
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	howmany := 0

	// Find all table cells with class "confluenceTd"
	doc.Find("td.confluenceTd").Each(func(index int, item *goquery.Selection) {
		// Print the text of each bill
		fmt.Println(item.Text())

		i++
		howmany++

		if i == 3 {
			fmt.Println() // Print a newline after every 3 bills
			i = 0
		}
		if howmany == 30 {
			return // Break the loop after 30 bills
		}
	})
}

