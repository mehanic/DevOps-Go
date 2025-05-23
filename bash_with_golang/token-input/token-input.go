package main

import (
	"fmt"
	"log"
	"github.com/jdkato/prose/v2"

)

func main() {
	// Input text
	inputText := "Here is some not very interesting text"

	// Create a new document
	doc, err := prose.NewDocument(inputText)
	if err != nil {
		log.Fatal(err)
	}

	// Tokenize the input text
	var tokens []string
	for _, tok := range doc.Tokens() {
		tokens = append(tokens, tok.Text)
	}

	// Print the tokens
	fmt.Println("Tokens:", tokens)
}

