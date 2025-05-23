package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// Define the structure for book and chapter
type Chapter struct {
	ChapterNumber string `xml:"chapterNumber"`
	ChapterTitle  string `xml:"chapterTitle"`
	PageCount     string `xml:"pageCount"`
}

type Book struct {
	Title           string    `xml:"title"`
	Publisher       string    `xml:"publisher"`
	NumberOfChapters string    `xml:"numberOfChapters"`
	PageCount       string    `xml:"pageCount"`
	Author          string    `xml:"author"`
	Chapters        []Chapter `xml:"chapters>chapter"`
}

// Root structure that holds all books
type Root struct {
	Books []Book `xml:"book"`
}

func parseBooks(filename string) {
	// Open the XML file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a struct to hold the parsed data
	var root Root

	// Parse the XML file
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&root); err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}

	// Print the books and their chapters
	for _, book := range root.Books {
		fmt.Printf("Book: %s, Publisher: %s, Chapters: %s, Pages: %s, Author: %s\n",
			book.Title, book.Publisher, book.NumberOfChapters, book.PageCount, book.Author)

		for _, chapter := range book.Chapters {
			fmt.Printf("  Chapter %s: %s, Pages: %s\n",
				chapter.ChapterNumber, chapter.ChapterTitle, chapter.PageCount)
		}
		fmt.Println() // Add a newline between books
	}
}

func main() {
	parseBooks("books.xml")
}
