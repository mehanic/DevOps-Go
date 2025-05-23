package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
	"strings"
)

// Function to fetch and print plain text content from a URL
func fetchTextContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Read the entire body as a string
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// Function to fetch and parse HTML content from a URL
func fetchHTMLContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Parse HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	// Find the content in <div id="mw-content-text">
	var contentBuilder strings.Builder
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "id" && attr.Val == "mw-content-text" {
					// Extract the text inside this <div>
					extractText(n, &contentBuilder)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return contentBuilder.String(), nil
}

// Helper function to extract text recursively from HTML nodes
func extractText(n *html.Node, contentBuilder *strings.Builder) {
	if n.Type == html.TextNode {
		contentBuilder.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, contentBuilder)
	}
}

func main() {
	// Fetch and print text content from a URL (text file)
	textURL := "http://www.pythonscraping.com/pages/warandpeace/chapter1-ru.txt"
	textContent, err := fetchTextContent(textURL)
	if err != nil {
		fmt.Println("Error fetching text content:", err)
		return
	}
	fmt.Println("Text Content:")
	fmt.Println(textContent)

	// Fetch and print HTML content from a URL (Wikipedia page)
	htmlURL := "http://en.wikipedia.org/wiki/Python_(programming_language)"
	htmlContent, err := fetchHTMLContent(htmlURL)
	if err != nil {
		fmt.Println("Error fetching HTML content:", err)
		return
	}
	fmt.Println("\nHTML Content:")
	fmt.Println(htmlContent)
}

