// find_images_in_page.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Make HTTP request
	response, err := http.Get("https://www.devdungeon.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and print image URLs
	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			fmt.Println(imgSrc)
		}
	})
}


// go run find-all-images-on-page.go                                                                                                                       
// https://www.devdungeon.com/sites/all/themes/devdungeon2/logo.png
// /sites/default/static/discord_join_dark.png
// /sites/all/themes/devdungeon2/icons/youtube-icon.png
// https://static.packt-cdn.com/products/9781788627917/cover/smaller
// /sites/default/static/UpporaAd1.png
