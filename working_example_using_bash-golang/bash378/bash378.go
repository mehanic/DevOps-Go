package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	// Equivalent to: echo "Hello world! This is my first Bash script!"
	fmt.Println("Hello world! This is my first Go script!")

	// Equivalent to: echo -n "I am executing the script with user: "
	fmt.Print("I am executing the script with user: ")

	// Equivalent to: whoami
	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("Error fetching user:", err)
	} else {
		fmt.Println(currentUser.Username)
	}

	// Equivalent to: echo -n "I am currently running in the directory: "
	fmt.Print("I am currently running in the directory: ")

	// Equivalent to: pwd
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error fetching current directory:", err)
	} else {
		fmt.Println(currentDir)
	}

	// Exit with code 0 (optional since Go exits with 0 by default)
	os.Exit(0)
}

