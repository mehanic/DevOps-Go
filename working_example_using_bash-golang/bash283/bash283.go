package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for MySQL user
	fmt.Print("MySQL User: ")
	userName, _ := reader.ReadString('\n')
	userName = userName[:len(userName)-1] // Trim newline

	// Prompt for MySQL password (silent input)
	fmt.Print("MySQL Password: ")
	mysqlPwd, _ := readPassword() // Read password silently
	fmt.Println() // Print a new line after password input

	// Prompt for MySQL command
	fmt.Print("MySQL Command: ")
	mysqlCmd, _ := reader.ReadString('\n')
	mysqlCmd = mysqlCmd[:len(mysqlCmd)-1] // Trim newline

	// Prompt for MySQL database
	fmt.Print("MySQL Database: ")
	mysqlDb, _ := reader.ReadString('\n')
	mysqlDb = mysqlDb[:len(mysqlDb)-1] // Trim newline

	// Prepare the MySQL command
	cmd := exec.Command("mysql", "-u", userName, "-p"+mysqlPwd, mysqlDb, "-e", mysqlCmd)

	// Set the command's output to the standard output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing MySQL command:", err)
	}
}

// readPassword reads a password from standard input without echoing.
func readPassword() (string, error) {
	var password string
	if _, err := fmt.Scanln(&password); err != nil {
		return "", err
	}
	return password, nil
}

