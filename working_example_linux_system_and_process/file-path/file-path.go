package main

import (
	"fmt"
	"os"
)

func main() {
	var reqPath string
	fmt.Print("Enter path to change working dir: ")
	fmt.Scan(&reqPath)

	// Print the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Println("The current working dir is: ", currentDir)

	// Change the working directory
	err = os.Chdir(reqPath)
	if err != nil {
		// Check specific errors
		if os.IsNotExist(err) {
			fmt.Println("Given path is not a valid path. So can't change working directory")
		} else if os.IsPermission(err) {
			fmt.Println("Sorry you don't have access for the given path. So can't change working directory")
		} else {
			fmt.Println("Given path is a file path. So can't change working directory")
		}
		return
	}

	// Print the new working directory
	newDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting new working directory:", err)
		return
	}
	fmt.Println("Now your new working dir is: ", newDir)
}

