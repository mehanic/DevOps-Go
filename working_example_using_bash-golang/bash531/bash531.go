package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

const ftpHome = "/var/ftp"

// RunCommand runs a system command with arguments and handles errors.
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Check if the script is run with root privileges.
func checkRoot() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	if currentUser.Uid != "0" {
		return fmt.Errorf("the script needs root privilege")
	}

	return nil
}

func main() {
	// Check for root privileges
	if err := checkRoot(); err != nil {
		fmt.Println("[*]", err)
		os.Exit(0)
	}

	// Install pure-ftpd using apt-get
	fmt.Println("Installing pure-ftpd...")
	if err := runCommand("apt-get", "install", "-y", "pure-ftpd"); err != nil {
		fmt.Println("Error installing pure-ftpd:", err)
		return
	}

	// Create FTP directory
	fmt.Println("Creating FTP directory:", ftpHome)
	if err := runCommand("mkdir", "-p", ftpHome); err != nil {
		fmt.Println("Error creating FTP directory:", err)
		return
	}

	// Add ftpgroup group
	fmt.Println("Adding ftpgroup group...")
	if err := runCommand("groupadd", "ftpgroup"); err != nil {
		fmt.Println("Error adding ftpgroup group:", err)
		return
	}

	// Add ftpuser user
	fmt.Println("Adding ftpuser user...")
	if err := runCommand("useradd", "-g", "ftpgroup", "-d", "/dev/null", "-s", "/etc", "ftpuser"); err != nil {
		fmt.Println("Error adding ftpuser user:", err)
		return
	}

	// Create ftpadmin user in Pure-FTPd
	fmt.Println("Adding ftpadmin user to Pure-FTPd...")
	if err := runCommand("pure-pw", "useradd", "ftpadmin", "-u", "ftpuser", "-d", ftpHome); err != nil {
		fmt.Println("Error adding ftpadmin user:", err)
		return
	}

	// Generate Pure-FTPd user database
	fmt.Println("Generating Pure-FTPd user database...")
	if err := runCommand("pure-pw", "mkdb"); err != nil {
		fmt.Println("Error generating Pure-FTPd user database:", err)
		return
	}

	// Create symlink for PureDB in /etc/pure-ftpd/auth/
	authPath := "/etc/pure-ftpd/auth"
	linkSrc := filepath.Join("..", "conf", "PureDB")
	linkDst := filepath.Join(authPath, "60pdb")
	fmt.Println("Creating symlink for PureDB in", authPath)
	if err := runCommand("ln", "-s", linkSrc, linkDst); err != nil {
		fmt.Println("Error creating symlink:", err)
		return
	}

	// Set permissions on FTP directory
	fmt.Println("Setting ownership for FTP directory...")
	if err := runCommand("chown", "-R", "ftpuser:ftpgroup", ftpHome); err != nil {
		fmt.Println("Error setting ownership:", err)
		return
	}

	// Start Pure-FTPd service
	fmt.Println("Starting Pure-FTPd service...")
	if err := runCommand("service", "pure-ftpd", "start"); err != nil {
		fmt.Println("Error starting Pure-FTPd service:", err)
		return
	}

	fmt.Println("FTP setup completed successfully.")
}

