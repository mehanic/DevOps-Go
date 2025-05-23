package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	numbers := []int{2, 3, 5, 7, 11, 13}

	// Open the file for writing
	filename := "numbers.json"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a JSON encoder and write to file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(numbers)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

