package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// URL to fetch CSV file
	url := "http://pythonscraping.com/files/MontyPythonAlbums.csv"

	// Fetch the CSV data from URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != 200 {
		log.Fatalf("Error: Status code %d while fetching URL", resp.StatusCode)
	}

	// Create a new CSV reader from the response body
	reader := csv.NewReader(resp.Body)

	// Read each row from the CSV
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV: %v", err)
		}

		// Print the CSV row
		fmt.Println(record)
	}
}

// [Name Year]
// [Monty Python's Flying Circus 1970]
// [Another Monty Python Record 1971]
// [Monty Python's Previous Record 1972]
// [The Monty Python Matching Tie and Handkerchief 1973]
// [Monty Python Live at Drury Lane 1974]
// [An Album of the Soundtrack of the Trailer of the Film of Monty Python and the Holy Grail 1975]
// [Monty Python Live at City Center 1977]
// [The Monty Python Instant Record Collection 1977]
// [Monty Python's Life of Brian 1979]
// [Monty Python's Cotractual Obligation Album 1980]
// [Monty Python's The Meaning of Life 1983]
// [The Final Rip Off 1987]
// [Monty Python Sings 1989]
// [The Ultimate Monty Python Rip Off 1994]
// [Monty Python Sings Again 2014]
