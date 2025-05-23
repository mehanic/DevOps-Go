package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Function to get the HTML node of a specific tag
func getNodeByTag(n *html.Node, tag atom.Atom) *html.Node {
	var crawler func(*html.Node) *html.Node
	crawler = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.DataAtom == tag {
			return n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if result := crawler(c); result != nil {
				return result
			}
		}
		return nil
	}
	return crawler(n)
}

// Function to extract text from an HTML node
func getTextFromNode(n *html.Node) string {
	if n == nil || n.FirstChild == nil {
		return ""
	}
	return n.FirstChild.Data
}

// Function to extract the text of the first paragraph within a node with a specific ID
func getFirstParagraphTextById(n *html.Node, id string) string {
	var crawler func(*html.Node) string
	crawler = func(n *html.Node) string {
		for _, attr := range n.Attr {
			if attr.Key == "id" && attr.Val == id {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					if c.DataAtom == atom.P && c.Type == html.ElementNode {
						return getTextFromNode(c)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if result := crawler(c); result != "" {
				return result
			}
		}
		return ""
	}
	return crawler(n)
}

func main() {
	// Step 1: Fetch the HTML page
	url := "https://www.debian.org/releases/stable/index.en.html"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch page: %v", err)
	}
	defer resp.Body.Close()

	// Parse the HTML document
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	// Step 2: Extract the title from the head element
	head := getNodeByTag(doc, atom.Head)
	titleNode := getNodeByTag(head, atom.Title)
	titleText := getTextFromNode(titleNode)

	// Step 3: Use regular expressions to find the release codename
	re := regexp.MustCompile(`\u201c(.*)\u201d`)
	matches := re.FindStringSubmatch(titleText)
	if len(matches) == 2 {
		release := matches[1]

		// Step 4: Extract the first paragraph from the content div
		content := getNodeByTag(doc, atom.Div)
		pText := getFirstParagraphTextById(content, "content")

		// Step 5: Extract version number from the paragraph text
		words := strings.Fields(pText)
		if len(words) > 1 {
			version := words[1]
			fmt.Printf("Codename: %s\nVersion: %s\n", release, version)
		} else {
			fmt.Println("Could not extract version")
		}
	} else {
		fmt.Println("Could not find codename in title")
	}
}

