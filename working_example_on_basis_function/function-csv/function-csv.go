package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func read(fileLocation string) ([][]string, error) {
	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '
	reader.Quote = '|'
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func write(fileLocation string, rows [][]string) error {
	file, err := os.Create(fileLocation)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ' '
	writer.Quote = '|'
	for _, row := range rows {
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

func rawTest() {
	var columns int
	fmt.Print("How many columns do you want to write? ")
	fmt.Scan(&columns)

	var inputRows [][]string
	keepGoing := true

	for keepGoing {
		var row []string
		for i := 0; i < columns; i++ {
			var input string
			fmt.Printf("column %d: ", i+1)
			fmt.Scan(&input)
			row = append(row, input)
		}
		inputRows = append(inputRows, row)

		var uiKeepGoing string
		fmt.Print("continue? (y/N): ")
		fmt.Scan(&uiKeepGoing)
		if strings.ToLower(uiKeepGoing) != "y" {
			keepGoing = false
		}
	}

	fmt.Println(inputRows)

	if err := write("raw.csv", inputRows); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	writtenValue, err := read("raw.csv")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println(writtenValue)
}

func main() {
	rawTest()
}

