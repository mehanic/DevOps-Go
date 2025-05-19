package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Global variable equivalent to FILES in Bash
	files := []string{"file1", "file2", "file3"}

	// Loop through each element in files
	for _, element := range files {
		createFile(element, "755") // Passing permissions as string
	}

	fmt.Println("Created all the files with a function!")
}

// createFile creates a file and sets its permissions
func createFile(fname string, permissions string) {
	// Create the file
	file, err := os.Create(fname)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after we're done

	// Change the permissions using chmod
	err = exec.Command("chmod", permissions, fname).Run()
	if err != nil {
		fmt.Println("Error changing permissions:", err)
		return
	}

	// List file details
	info, err := os.Stat(fname)
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	// Print file details (similar to `ls -l`)
	fmt.Printf("%s\n", info.Mode().Perm())
}

