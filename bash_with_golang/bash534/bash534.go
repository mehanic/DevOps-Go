package main

import (
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
)

// Utility function to run shell commands
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Step 1: Install atftpd package using apt-get
	fmt.Println("Installing atftpd package...")
	if err := runCommand("apt-get", "install", "-y", "atftpd"); err != nil {
		fmt.Println("Error installing atftpd:", err)
		return
	}

	// Step 2: Create /srv/tftp directory
	fmt.Println("Creating /srv/tftp directory...")
	if err := os.MkdirAll("/srv/tftp", 0755); err != nil {
		fmt.Println("Error creating /srv/tftp directory:", err)
		return
	}

	// Step 3: Change ownership of /srv/tftp to nobody
	fmt.Println("Changing ownership of /srv/tftp to nobody...")
	if err := runCommand("chown", "-R", "nobody", "/srv/tftp"); err != nil {
		fmt.Println("Error changing ownership of /srv/tftp:", err)
		return
	}

	// Step 4: Create a sample TFTP file (tftp.txt) with content "hello tftp-server"
	fmt.Println("Creating sample TFTP file /srv/tftp/tftp.txt...")
	tftpContent := "hello tftp-server\n"
	if err := ioutil.WriteFile("/srv/tftp/tftp.txt", []byte(tftpContent), 0644); err != nil {
		fmt.Println("Error creating /srv/tftp/tftp.txt:", err)
		return
	}

	// Step 5: Start the atftpd service
	fmt.Println("Starting atftpd service...")
	if err := runCommand("service", "atftpd", "start"); err != nil {
		fmt.Println("Error starting atftpd service:", err)
		return
	}

	fmt.Println("TFTP server setup completed successfully.")
}

