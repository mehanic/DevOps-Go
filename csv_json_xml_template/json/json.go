package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// Read a JSON file
	reqFile := "myjson.json"

	// Open the JSON file
	fo, err := os.Open(reqFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fo.Close()

	// Read the file content
	data, err := io.ReadAll(fo)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the file is empty
	if len(data) == 0 {
		fmt.Println("Error: JSON file is empty")
		return
	}

	// Unmarshal the JSON data
	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the JSON data
	fmt.Printf("Parsed JSON: %+v\n", result)

	// Write data (dictionary data) into a JSON file
	myDict := map[string]interface{}{
		"Name":   "Narendra",
		"skills": []string{"Python", "shell", "yaml", "AWS"},
	}

	reqFile = "myinfo.json"

	// Create or open the JSON file
	fo, err = os.Create(reqFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fo.Close()

	// Marshal the dictionary data to JSON
	jsonData, err := json.MarshalIndent(myDict, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the JSON data to the file
	_, err = fo.Write(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Parsed JSON: map[Age:30 Name:John Doe Skills:[Go Python Docker]]
