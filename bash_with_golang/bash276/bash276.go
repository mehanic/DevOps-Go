package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
//	"strings"
)

func main() {
	// Define log file and output CSV file
	gcLog := "server-myapp-gc.log"
	csvLoc := "/tmp/gclog.csv"

	// Create the output CSV file
	csvFile, err := os.Create(csvLoc)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write header to CSV
	writer.Write([]string{"Heap Usage Before", "Heap Usage After"})

	// Open the GC log file
	file, err := os.Open(gcLog)
	if err != nil {
		fmt.Println("Error opening GC log file:", err)
		return
	}
	defer file.Close()

	// Compile the regular expression to match the relevant line
	re := regexp.MustCompile(`(?i)Full.*?(\d+)K->(\d+)K`)

	// Read through the log file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			// Extract heap usage information using regex
			matches := re.FindStringSubmatch(line)
			if len(matches) == 3 {
				before := matches[1]
				after := matches[2]

				// Write to the CSV file
				writer.Write([]string{before, after})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading GC log file:", err)
		return
	}

	fmt.Printf("GC log mapping completed. Please copy the %s file and open it in Excel to view the GC map.\n", csvLoc)
}

