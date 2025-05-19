package main

import (
	"fmt"
	"os"
	"os/user"
)

func isUserRoot() bool {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return false
	}
	
	// Check if the user ID is "0" (root)
	return currentUser.Uid == "0"
}

func main() {
	if isUserRoot() {
		fmt.Println("You are root user, you can go ahead.")
	} else {
		fmt.Println("You need to be administrator to run this script.")
		os.Exit(1) // Exit with status 1 if not root
	}
}

