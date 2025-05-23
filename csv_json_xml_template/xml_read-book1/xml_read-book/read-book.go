package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Define the structure that reflects the XML elements
type Book struct {
	XMLName xml.Name `xml:"book"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
}

// Root structure for the list of books
type Books struct {
	XMLName xml.Name `xml:"root"`
	Books   []Book   `xml:"book"`
}

func main() {
	// Open the XML file
	file, err := os.Open("books.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a struct to hold the data
	var books Books

	// Parse the XML file
	if err := xml.NewDecoder(file).Decode(&books); err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	// Print root tag
	fmt.Println("Root tag:", books.XMLName.Local)

	// Access attributes and child elements
	for _, book := range books.Books {
		fmt.Printf("Book ID: %s, Name: %s\n", book.ID, book.Name)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
	}
}


// Root tag: root
// Book ID: book1, Name: Learning Python 2
// Title: Learning Python 2
// Author: Author1
// Book ID: book2, Name: Learning Python 3
// Title: Learning Python 3
// Author: Author2
