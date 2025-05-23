package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Specify the path to the scripts directory
	scriptsDir := "/path/to/scripts/dir"

	// Read the contents of the directory
	files, err := ioutil.ReadDir(scriptsDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Iterate over each file in the directory
	for _, file := range files {
		// Build the full path to the script
		scriptPath := filepath.Join(scriptsDir, file.Name())

		// Check if the file is a regular file and executable
		if file.Mode().IsRegular() && file.Mode()&0111 != 0 {
			// Execute the script
			fmt.Printf("Executing script: %s\n", scriptPath)
			cmd := exec.Command(scriptPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			// Run the command and check for errors
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error executing script %s: %v\n", scriptPath, err)
			}
		}
	}
}

