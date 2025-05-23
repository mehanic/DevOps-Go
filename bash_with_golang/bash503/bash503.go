package main

import (
	"fmt"
	"os"
	"os/exec"
//	"strconv"
	"strings"
)

const newUserName = "bob"
const defaultPassword = "password"

func main() {
	// Verify this program is run with root privileges.
	if os.Geteuid() != 0 {
		fmt.Println("Please run as root or with sudo!")
		os.Exit(1)
	}

	// Check if the user exists.
	if !userExists(newUserName) {
		// User does not exist, create the user.
		if err := createUser(newUserName); err != nil {
			fmt.Println("Error creating user:", err)
			os.Exit(1)
		}
	}

	// Set the password for the user.
	if err := setPassword(newUserName, defaultPassword); err != nil {
		fmt.Println("Error setting password:", err)
		os.Exit(1)
	}

	fmt.Println("Password set successfully.")
}

// userExists checks if a user exists on the system.
func userExists(username string) bool {
	cmd := exec.Command("id", username)
	err := cmd.Run()
	return err == nil
}

// createUser creates a new user with the specified username.
func createUser(username string) error {
	cmd := exec.Command("useradd", "-m", username)
	return cmd.Run()
}

// setPassword sets the password for the specified user.
func setPassword(username, password string) error {
	// Prepare the input for chpasswd.
	input := fmt.Sprintf("%s:%s", username, password)
	cmd := exec.Command("chpasswd")
	cmd.Stdin = strings.NewReader(input)
	return cmd.Run()
}

