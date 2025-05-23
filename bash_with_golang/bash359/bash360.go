package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "File2"

	// Check if the file exists
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	}

	// Check if the file has read permission
	if fileInfo, err := os.Stat(fileName); err == nil {
		if fileInfo.Mode().Perm()&0400 != 0 { // Read permission
			fmt.Println("File has read permission")
			fmt.Println(0) // Equivalent to success exit code
		} else {
			fmt.Println("File does not have read permission")
			fmt.Println(1) // Equivalent to failure exit code
		}
	}

	// Check if the file has write permission
	if fileInfo, err := os.Stat(fileName); err == nil {
		if fileInfo.Mode().Perm()&0200 != 0 { // Write permission
			fmt.Println("File has write permission")
			fmt.Println(0) // Equivalent to success exit code
		} else {
			fmt.Println("File does not have write permission")
			fmt.Println(1) // Equivalent to failure exit code
		}
	}

	// Check if the file has execute permission
	if fileInfo, err := os.Stat(fileName); err == nil {
		if fileInfo.Mode().Perm()&0100 != 0 { // Execute permission
			fmt.Println("File has execute permission")
			fmt.Println(0) // Equivalent to success exit code
		} else {
			fmt.Println("File does not have execute permission")
			fmt.Println(1) // Equivalent to failure exit code
		}
	}
}

