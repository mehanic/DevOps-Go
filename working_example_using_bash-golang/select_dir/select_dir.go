package main

import (
	"fmt"
	"io/ioutil"
//	"os"
	"path/filepath"
)

func main() {
	// Step 1: List all directories in the root directory
	files, err := ioutil.ReadDir("/")
	if err != nil {
		fmt.Println("Error reading root directory:", err)
		return
	}

	// Step 2: Build the directory list, including the "Finished" option
	directoryList := []string{"Finished"}
	for _, file := range files {
		if file.IsDir() {
			directoryList = append(directoryList, filepath.Join("/", file.Name()))
		}
	}

	// Step 3: Initialize user selection
	var directory string

	// Step 4: Loop until the user selects "Finished"
	for {
		fmt.Println("\nSelect a directory to process:")
		for i, dir := range directoryList {
			fmt.Printf("%d) %s\n", i+1, dir) // Display directories with index
		}

		// Step 5: Get user's choice
		var choice int
		fmt.Print("Directory to process? ")
		_, err := fmt.Scan(&choice)
		if err != nil || choice < 1 || choice > len(directoryList) {
			fmt.Println("Invalid selection!")
			continue
		}

		// Step 6: Handle the user's selection
		directory = directoryList[choice-1]
		if directory == "Finished" {
			fmt.Println("Finished processing directories.")
			break
		}

		fmt.Printf("You chose number %d, processing %s...\n", choice, directory)
		// Here, you can perform any desired processing on the selected directory
	}
}

