package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Define the filename
	filename := "username.json"

	// Open the JSON file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Declare a variable to hold the username
	var username string

	// Decode the JSON file into the variable
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&username)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Print the welcome message
	fmt.Printf("Welcome back, %s!\n", username)
}

