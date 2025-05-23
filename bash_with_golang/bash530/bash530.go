package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Run a system command and check for errors
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Check if a command exists (equivalent of `which` in bash)
func commandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func main() {
	// Install packages using apt-get
	fmt.Println("Installing apache2, libapache2-mod-php, and phpmyadmin...")
	if err := runCommand("apt-get", "install", "-y", "apache2", "libapache2-mod-php", "phpmyadmin"); err != nil {
		fmt.Println("Error installing packages:", err)
		return
	}

	// Check if mysqld exists
	if !commandExists("mysqld") {
		fmt.Println("Please install mysql-server first")
		return
	}

	// Start MySQL service
	fmt.Println("Starting MySQL service...")
	if err := runCommand("service", "mysql", "start"); err != nil {
		fmt.Println("Error starting MySQL service:", err)
		return
	}

	// Start Apache2 service
	fmt.Println("Starting Apache2 service...")
	if err := runCommand("service", "apache2", "start"); err != nil {
		fmt.Println("Error starting Apache2 service:", err)
		return
	}

	fmt.Println("Services started successfully.")
}

