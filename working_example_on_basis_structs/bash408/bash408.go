package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
//	"os"
)

// Struct to hold the XML data
type Data struct {
	XMLName xml.Name `xml:"root"` // Adjust according to your XML structure
	Items   []Item   `xml:"item"` // Adjust according to your XML structure
}

type Item struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
	// Add more fields based on your XML structure
}

// Convert XML to JSON
func xmlToJSON(inputFile string) (string, error) {
	xmlData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return "", err
	}

	var data Data
	if err := xml.Unmarshal(xmlData, &data); err != nil {
		return "", err
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

// Write JSON to file
func writeJSONToFile(outputFile, jsonData string) error {
	return ioutil.WriteFile(outputFile, []byte(jsonData), 0644)
}

func main() {
	inputFile := "words.xml"
	outputFile := "words.json"

	// Convert XML to JSON
	jsonData, err := xmlToJSON(inputFile)
	if err != nil {
		log.Fatalf("Error converting XML to JSON: %v", err)
	}

	// Write JSON data to file
	if err := writeJSONToFile(outputFile, jsonData); err != nil {
		log.Fatalf("Error writing JSON to file: %v", err)
	}

	fmt.Println("Conversion successful! JSON written to", outputFile)
}

