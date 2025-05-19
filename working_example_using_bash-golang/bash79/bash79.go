package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watchDir := filepath.Join(os.Getenv("HOME"), "Desktop", "abc")
	oldDir := filepath.Join(os.Getenv("HOME"), "Desktop", "Old_abc")
	outputFile := "output.txt"

	// Create a new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Set up a folder to watch
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				// If a file is created
				if event.Op&fsnotify.Create == fsnotify.Create {
					fmt.Println("File created:", event.Name)

					// Get the current date and time for renaming
					cdate := time.Now().Format("2006-01-02-15:04")

					// Move the file output.txt from abc folder to Old_abc folder with the new name
					src := filepath.Join(watchDir, outputFile)
					dest := filepath.Join(oldDir, fmt.Sprintf("%s-%s", cdate, outputFile))

					if err := moveFile(src, dest); err != nil {
						log.Printf("Error moving file: %v\n", err)
					} else {
						fmt.Printf("Moved %s to %s\n", src, dest)
					}
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	// Start watching the directory
	err = watcher.Add(watchDir)
	if err != nil {
		log.Fatal(err)
	}

	// Block forever
	<-make(chan struct{})
}

// moveFile moves a file from source to destination
func moveFile(src, dest string) error {
	// Use the mv system command to move the file
	cmd := exec.Command("mv", src, dest)
	return cmd.Run()
}

