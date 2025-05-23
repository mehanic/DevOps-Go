package main

import (
	"encoding/csv"
	"fmt"
	"log"
//	"os"
	"strings"
)

var logRows = [][]string{
	{"date", "engine on", "fuel height"},
	{"", "engine off", "fuel height"},
	{"", "Other notes", ""},
	{"", "", ""},
	{"10/25/13", "08:24:00 AM", "29"},
	{"", "01:15:00 PM", "27"},
	{"", "calm seas -- anchor solomon's island", ""},
	{"10/26/13", "09:12:00 AM", "27"},
	{"", "06:25:00 PM", "22"},
	{"", "choppy -- anchor in jackson's creek", ""},
}

func main() {
	// Example for slicing head and tail of logRows
	head := logRows[:4]
	tail := logRows[4:]

	// Step 1: Head and Tail
	fmt.Println("Step 1: Head and Tail")
	fmt.Println("Head:", head)
	fmt.Println("Tail:", tail)

	// Step 2: Extract every 3rd row from the tail, starting from index 2
	fmt.Println("\nStep 2: Every 3rd row from the tail starting at index 2")
	for i := 2; i < len(tail); i += 3 {
		fmt.Println(tail[i])
	}

	// Step 3: Zipping pairs of every 3rd row in the tail
	fmt.Println("\nStep 3: Zipping pairs of every 3rd row")
	for i := 0; i < len(tail); i += 3 {
		if i+1 < len(tail) {
			fmt.Println(tail[i], tail[i+1])
		}
	}

	// Reading CSV file (mocked inline data for example)
	fmt.Println("\nReading CSV:")
	reader := csv.NewReader(strings.NewReader(mockCSVData()))
	logRowsFromFile, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("First row:", logRowsFromFile[0])
	fmt.Println("Last row:", logRowsFromFile[len(logRowsFromFile)-1])
	fmt.Println("Full content:", logRowsFromFile)
}

// Mock CSV data for reading (this would be a file in actual implementation)
func mockCSVData() string {
	return `
date,engine on,fuel height
,engine off,fuel height
,Other notes,
,,
10/25/13,08:24:00 AM,29
,01:15:00 PM,27
,calm seas -- anchor solomon's island,
10/26/13,09:12:00 AM,27
,06:25:00 PM,22
,choppy -- anchor in jackson's creek,
`
}


// Step 1: Head and Tail
// Head: [[date engine on fuel height] [ engine off fuel height] [ Other notes ] [  ]]
// Tail: [[10/25/13 08:24:00 AM 29] [ 01:15:00 PM 27] [ calm seas -- anchor solomon's island ] [10/26/13 09:12:00 AM 27] [ 06:25:00 PM 22] [ choppy -- anchor in jackson's creek ]]

// Step 2: Every 3rd row from the tail starting at index 2
// [ calm seas -- anchor solomon's island ]
// [ choppy -- anchor in jackson's creek ]

// Step 3: Zipping pairs of every 3rd row
// [10/25/13 08:24:00 AM 29] [ 01:15:00 PM 27]
// [10/26/13 09:12:00 AM 27] [ 06:25:00 PM 22]

// Reading CSV:
// First row: [date engine on fuel height]
// Last row: [ choppy -- anchor in jackson's creek ]
// Full content: [[date engine on fuel height] [ engine off fuel height] [ Other notes ] [  ] [10/25/13 08:24:00 AM 29] [ 01:15:00 PM 27] [ calm seas -- anchor solomon's island ] [10/26/13 09:12:00 AM 27] [ 06:25:00 PM 22] [ choppy -- anchor in jackson's creek ]]
