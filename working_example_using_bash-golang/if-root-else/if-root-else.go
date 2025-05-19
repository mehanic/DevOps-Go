package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	// Get current user information
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error retrieving current user:", err)
		os.Exit(1)
	}

	// Check if UID is 0 (root user)
	if currentUser.Uid != "0" {
		fmt.Println("Non root user. Please run as root.")
	} else {
		fmt.Println("Root user")
	}
}

