package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open the JSON file
	file, err := os.Open("books.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Define a variable to hold the JSON data
	var books interface{}

	// Decode JSON data into the variable
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&books)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the JSON data
	fmt.Println(books)

	// Print the type of the data
	fmt.Printf("Type: %T\n", books)
}

// [map[author:author chapters:[map[chapterNumber:1 chapterTitle:Python Fundamentals pageCount:30] map[chapterNumber:2 chapterTitle:Chapter 2 pageCount:25]] numberOfChapters:12 pageCount:500 publisher:Packt Publishing title:Learning Python 3] map[author:author chapters:[map[chapterNumber:1 chapterTitle:Python Fundamentals pageCount:30] map[chapterNumber:2 chapterTitle:Chapter 2 pageCount:25]] numberOfChapters:12 pageCount:500 publisher:Packt Publishing title:Learning Python 3]]
// Type: []interface {}
