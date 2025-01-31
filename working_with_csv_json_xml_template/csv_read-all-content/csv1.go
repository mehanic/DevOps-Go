package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	reqFile := "new_info.csv"

	// Open the file
	fo, err := os.Open(reqFile)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer fo.Close()

	// Create a new CSV reader with '|' as the delimiter
	reader := csv.NewReader(fo)
	reader.Comma = '|'

	// Read all content
	content, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV content: %s", err)
	}

	// Uncomment the following lines to print the content or header
	// fmt.Println(content)
	// fmt.Printf("The header is:\n %v\n", content[0])
	// header := content[0]
	// fmt.Println("The header is: ", header)

	// Print the number of rows excluding the header
	fmt.Println("The number of rows are: ", len(content)-1)

	// Uncomment the following lines to iterate over each row
	/*
	   for _, each := range content {
	       fmt.Println(each)
	   }
	*/
}

//The number of rows are:  3
