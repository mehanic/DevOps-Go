package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
//	"path/filepath"
	"syscall"
)

// Check and print the permissions of the file or directory
func permissions(path string) {
	fmt.Printf("\nWhat are our permissions on this %s?\n", path)
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error accessing permissions: %v\n", err)
		return
	}

	// Check permissions
	if info.Mode().Perm()&0400 != 0 {
		fmt.Println("[R] Read")
	}
	if info.Mode().Perm()&0200 != 0 {
		fmt.Println("[W] Write")
	}
	if info.Mode().Perm()&0100 != 0 {
		fmt.Println("[X] Exec")
	}
}

// Check and print the attributes of the file
func fileAttributes(path string) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error retrieving file attributes: %v\n", err)
		return
	}

	// Check if the file is empty
	if info.Size() == 0 {
		fmt.Printf("\"%s\" is empty\n", path)
		return
	}

	// Get file size
	fmt.Printf("\"%s\" file size is: %d bytes\n", path, info.Size())

	// Get current user
	user, err := user.Current()
	if err != nil {
		fmt.Printf("Error retrieving user info: %v\n", err)
		return
	}

	// Check ownership
	if fileInfo, ok := info.Sys().(*syscall.Stat_t); ok {
		if user.Uid != fmt.Sprint(fileInfo.Uid) {
			fmt.Printf("%s is not the owner of \"%s\"\n", user.Username, path)
		}
		if user.Gid != fmt.Sprint(fileInfo.Gid) {
			fmt.Printf("%s is not among the owning group(s) for \"%s\"\n", user.Username, path)
		}
	}

	// Print permissions
	permissions(path)
}

// Check and print the attributes of the directory
func dirAttributes(path string) {
	fmt.Printf("Directory \"%s\" has children:\n\n", path)
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	// Print permissions
	permissions(path)
}

// Main function to handle file checkout
func checkoutFile() {
	var fileToTest string
	fmt.Print("\nWhat is the complete path of the file you want to inspect?\n# ")
	fmt.Scanln(&fileToTest)

	// Check if the path exists
	if _, err := os.Stat(fileToTest); os.IsNotExist(err) {
		fmt.Printf("Error: \"%s\" does not exist!\n", fileToTest)
		return
	}

	if info, err := os.Stat(fileToTest); err == nil {
		if info.IsDir() {
			dirAttributes(fileToTest)
		} else {
			fileAttributes(fileToTest)
		}
		checkoutFile() // Recursively checkout file/directory
	}
}

func main() {
	fmt.Println("Welcome to the file attributes tester")
	fmt.Println("To exit, press CTRL + C")

	checkoutFile()
}

