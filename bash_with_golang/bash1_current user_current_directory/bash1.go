package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	// Print a welcome message
	fmt.Println("Hello world! This is my first Go program!")

	// Get current user
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}
	fmt.Print("I am executing the script with user: ")
	fmt.Println(currentUser.Username)

	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Print("I am currently running in the directory: ")
	fmt.Println(wd)

	// Exit with a success status
	os.Exit(0)
}

