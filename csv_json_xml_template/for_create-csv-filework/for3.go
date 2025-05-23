package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Open the CSV file for reading
	file, err := os.Open("my.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read and print each record from the CSV file
	for {
		record, err := reader.Read()
		if err != nil {
			// Break the loop if the end of the file is reached or an error occurs
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading record:", err)
			return
		}
		fmt.Println(record)
	}
}

