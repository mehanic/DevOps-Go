package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// Define the directory and source file
	dirName := "dpkg"
	sourceFile := "/var/log/dpkg.log"

	// Create the directory if it doesn't exist
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err := os.Mkdir(dirName, 0755); err != nil {
			fmt.Println("Cannot create the directory, stopping script.")
			os.Exit(1)
		}
	}

	// Define the destination file path
	destinationFile := filepath.Join(dirName, "dpkg.log")

	// Read the source file
	data, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println("Cannot read dpkg.log:", err)
		os.Exit(1)
	}

	// Write the data to the destination file
	if err := ioutil.WriteFile(destinationFile, data, 0644); err != nil {
		fmt.Println("Cannot copy dpkg.log to the new directory:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully copied dpkg.log to the dpkg directory.")
}

