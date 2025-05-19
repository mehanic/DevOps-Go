package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/unidoc/unioffice/document"
)

// getText extracts text from the given .docx file
func getText(filename string) (string, error) {
	// Open the .docx file
	doc, err := document.Open(filename)
	if err != nil {
		return "", err
	}
	defer doc.Close()

	var fullText []string

	// Loop through the paragraphs and extract text
	for _, para := range doc.Paragraphs() {
		fullText = append(fullText, para.Text())
	}

	// Join the paragraphs with double newlines, similar to Python's '\n\n'
	return strings.Join(fullText, "\n\n"), nil
}

func main() {
	// Replace 'your-file.docx' with the actual filename
	filename := "your-file.docx"

	// Call the getText function
	text, err := getText(filename)
	if err != nil {
		log.Fatalf("Failed to extract text: %v", err)
	}

	// Print the extracted text
	fmt.Println(text)
}

