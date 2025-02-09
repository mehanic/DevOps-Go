package main

import (
	"fmt"
	"strings"
	"time"
)

// Article represents the structure of the article
type Article struct {
	URL        string
	Title      string
	Text       []string
	LastUpdated string
}

// WikispiderPipeline processes the scraped article data
type WikispiderPipeline struct{}

// processItem processes an Article and returns the modified Article
func (p *WikispiderPipeline) processItem(article Article) Article {
	// Remove the phrase from LastUpdated
	article.LastUpdated = strings.ReplaceAll(article.LastUpdated, "This page was last edited on", "")
	article.LastUpdated = strings.TrimSpace(article.LastUpdated)

	// Parse the LastUpdated date
	layout := "2 January 2006, at 15:04."
	if t, err := time.Parse(layout, article.LastUpdated); err == nil {
		// Format it back to a string if needed
		article.LastUpdated = t.Format(time.RFC3339) // Example formatting, can be adjusted
	} else {
		fmt.Println("Error parsing date:", err)
	}

	// Remove empty lines from Text
	var nonEmptyLines []string
	for _, line := range article.Text {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	// Join non-empty lines into a single string
	article.Text = nonEmptyLines
	return article
}

// Example usage
func main() {
	pipeline := &WikispiderPipeline{}

	// Example article data
	article := Article{
		URL:        "https://en.wikipedia.org/wiki/Benevolent_dictator_for_life",
		Title:      "Benevolent Dictator for Life",
		Text:       []string{"This is a line of text.", " ", "\n", "Another line of text."},
		LastUpdated: "This page was last edited on 15 February 2024, at 12:30.",
	}

	// Process the article
	processedArticle := pipeline.processItem(article)

	// Output processed article
	fmt.Printf("Processed Article:\n")
	fmt.Printf("URL: %s\n", processedArticle.URL)
	fmt.Printf("Title: %s\n", processedArticle.Title)
	fmt.Printf("Text: %s\n", strings.Join(processedArticle.Text, "\n"))
	fmt.Printf("Last Updated: %s\n", processedArticle.LastUpdated)
}

// Processed Article:
// URL: https://en.wikipedia.org/wiki/Benevolent_dictator_for_life
// Title: Benevolent Dictator for Life
// Text: This is a line of text.
// Another line of text.
// Last Updated: 2024-02-15T12:30:00Z
