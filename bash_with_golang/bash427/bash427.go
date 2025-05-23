package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"gopkg.in/fsnotify.v1"
)

func main() {
	// Define the folder to watch and the old output destination
	folder := filepath.Join(os.Getenv("HOME"), "Desktop", "abc")
	oldOutputDir := filepath.Join(os.Getenv("HOME"), "Desktop", "Old_abc")

	// Create the old output directory if it doesn't exist
	if err := os.MkdirAll(oldOutputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create old output directory: %v", err)
	}

	// Create a new watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}
	defer watcher.Close()

	// Start watching the specified folder
	if err := watcher.Add(folder); err != nil {
		log.Fatalf("Failed to add folder to watcher: %v", err)
	}

	// Start monitoring for events
	fmt.Printf("Monitoring %s for new files...\n", folder)
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			// Check if the event is a creation event
			if event.Op&fsnotify.Create == fsnotify.Create {
				// Get the current date and time for the new filename
				cdate := time.Now().Format("2006-01-02-15:04")

				// Define the output file path
				outputFile := filepath.Join(folder, "output.txt")
				newOutputFile := filepath.Join(oldOutputDir, fmt.Sprintf("%s-output.txt", cdate))

				// Move the output.txt file to the new location
				if _, err := os.Stat(outputFile); err == nil {
					if err := os.Rename(outputFile, newOutputFile); err != nil {
						log.Printf("Failed to move %s to %s: %v", outputFile, newOutputFile, err)
					} else {
						fmt.Printf("Moved %s to %s\n", outputFile, newOutputFile)
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Error: %v", err)
		}
	}
}

