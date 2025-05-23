package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Function to create a file and set permissions
func createFile(fname string, permissions os.FileMode) error {
	// Create the file
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, permissions)
	if err != nil {
		return err
	}
	defer file.Close() // Ensure the file is closed after operations

	// Use the `chmod` command to set permissions (in case it's not possible in os.FileMode)
	cmd := exec.Command("chmod", fmt.Sprintf("%o", permissions.Perm()), fname)
	if err := cmd.Run(); err != nil {
		return err
	}

	// List the file details
	cmd = exec.Command("ls", "-l", fname)
	cmd.Stdout = os.Stdout // Redirect output to standard output
	return cmd.Run()
}

func main() {
	files := []string{"file1", "file2", "file3"} // This is a global variable

	for _, element := range files {
		if err := createFile(element, 0755); err != nil { // 0755 is a+x
			fmt.Printf("Error creating file %s: %v\n", element, err)
		}
	}

	fmt.Println("Created all the files with a function!")
}

