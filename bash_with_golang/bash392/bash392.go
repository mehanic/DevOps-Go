package main

import (
	"fmt"
	"os"
	"os/user"
)

var globalRet = 255

// myFunctionGlobal lists the .bashrc file of the current user
func myFunctionGlobal() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return
	}

	// Attempt to list the .bashrc file
	_, err = os.Stat("/home/" + usr.Username + "/.bashrc")
	if err == nil {
		fmt.Println("Found .bashrc at /home/" + usr.Username + "/.bashrc")
		globalRet = 0 // Update globalRet on success
	} else if os.IsNotExist(err) {
		fmt.Println("The file /home/" + usr.Username + "/.bashrc does not exist.")
		globalRet = 1 // Update globalRet for file not found
	} else {
		fmt.Println("Error checking for .bashrc:", err)
		globalRet = 255 // Update globalRet for other errors
	}
}

// myFunctionReturn lists the .bashrc file of the current user and returns an error code
func myFunctionReturn() int {
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return 255
	}

	// Attempt to list the .bashrc file
	_, err = os.Stat("/home/" + usr.Username + "/.bashrc")
	if err == nil {
		fmt.Println("Found .bashrc at /home/" + usr.Username + "/.bashrc")
		return 0 // Return 0 on success
	} else if os.IsNotExist(err) {
		fmt.Println("The file /home/" + usr.Username + "/.bashrc does not exist.")
		return 1 // Return 1 for file not found
	} else {
		fmt.Println("Error checking for .bashrc:", err)
		return 255 // Return 255 for other errors
	}
}

// myFunctionStr checks for the existence of a user's .bashrc file and returns a message
func myFunctionStr(uname string) string {
	output := ""
	if _, err := os.Stat("/home/" + uname + "/.bashrc"); err == nil {
		output = "FOUND IT"
	} else {
		output = "NOT FOUND"
	}
	return output
}

func main() {
	// Display the current global return value
	fmt.Println("Current ret:", globalRet)

	// Call the first function and check for .bashrc
	myFunctionGlobal()
	fmt.Println("Current ret after myFunctionGlobal:", globalRet)

	// Reset global return value
	globalRet = 255
	fmt.Println("Current ret:", globalRet)

	// Call the second function and check for .bashrc
	globalRet = myFunctionReturn()
	fmt.Println("Current ret after myFunctionReturn:", globalRet)

	// Example usage of myFunctionStr
	usr, err := user.Current()
	if err == nil {
		result := myFunctionStr(usr.Username)
		fmt.Println(result)
	} else {
		fmt.Println("Error getting current user:", err)
	}
}

