package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// Step 1: Fetch the HTML page
	resp, err := http.Get("http://www.pythonscraping.com")
	if err != nil {
		fmt.Println("Error fetching the webpage:", err)
		return
	}
	defer resp.Body.Close()

	// Step 2: Parse the HTML using the net/html package
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// Step 3: Find the image URL (based on the "a" tag with id="logo")
	var imageURL string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == "logo" {
					// Now find the child img tag
					for c := n.FirstChild; c != nil; c = c.NextSibling {
						if c.Type == html.ElementNode && c.Data == "img" {
							for _, imgAttr := range c.Attr {
								if imgAttr.Key == "src" {
									imageURL = imgAttr.Val
									return
								}
							}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	if imageURL == "" {
		fmt.Println("Image URL not found")
		return
	}

	// Step 4: Download the image
	if !strings.HasPrefix(imageURL, "http") {
		imageURL = "http://www.pythonscraping.com" + imageURL
	}

	resp, err = http.Get(imageURL)
	if err != nil {
		fmt.Println("Error downloading the image:", err)
		return
	}
	defer resp.Body.Close()

	// Create the file
	file, err := os.Create("logo.jpg")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()

	// Copy the image content to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return
	}

	fmt.Println("Image successfully downloaded and saved as logo.jpg")
}

