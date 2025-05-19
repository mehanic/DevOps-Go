package main

import "fmt"

func main() {
	// Define the MY_SHELL variable
	myShell := "csh"

	// Check if MY_SHELL is "bash"
	if myShell == "bash" {
		fmt.Println("You seem to like the bash shell.")
	} else {
		fmt.Println("You don't seem to like the bash shell.")
	}
}

