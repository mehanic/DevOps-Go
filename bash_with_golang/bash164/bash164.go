package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// Get the current date in YYYY-MM-DD format
	currentDate := time.Now().Format("2006-01-02")

	// Use filepath.Glob to list all .jpg files in the current directory
	pictures, err := filepath.Glob("*.jpg")
	if err != nil {
		log.Fatalf("Failed to list jpg files: %v", err)
	}

	// Loop through each picture and rename it
	for _, picture := range pictures {
		// Construct the new file name by prepending the date
		newName := fmt.Sprintf("%s-%s", currentDate, picture)
		fmt.Printf("Renaming %s to %s\n", picture, newName)

		// Rename the file
		err := os.Rename(picture, newName)
		if err != nil {
			log.Printf("Failed to rename %s: %v", picture, err)
		}
	}
}

