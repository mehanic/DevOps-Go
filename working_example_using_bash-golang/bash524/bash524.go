package main

import (
	"fmt"
	"os"
//	"os/exec"
)

// Function to create a directory and change into it
func mcd(dirName string) error {
	// Create the directory
	err := os.Mkdir(dirName, 0755) // Set permissions for the directory
	if err != nil {
		return err // Return error if the directory cannot be created
	}

	// Change the working directory (only within the Go program)
	err = os.Chdir(dirName)
	if err != nil {
		return err // Return error if the directory cannot be changed into
	}

	return nil
}

func main() {
	dirName := "test1" // The directory name you want to create

	// Call the mcd function
	err := mcd(dirName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the current working directory
	currentDir, _ := os.Getwd() // Get the current working directory
	fmt.Println("Current directory:", currentDir)

	// Optionally, you can run a command in the new directory
	// Uncomment the following lines to execute a command
	// cmd := exec.Command("ls") // For example, list files in the new directory
	// cmd.Dir = currentDir       // Set the command's working directory
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	//     fmt.Println("Error executing command:", err)
	// } else {
	//     fmt.Println(string(output))
	// }
}

