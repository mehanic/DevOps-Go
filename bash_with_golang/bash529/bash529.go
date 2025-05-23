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

func main() {
	// Install dnsmasq using apt-get
	fmt.Println("Installing dnsmasq...")
	if err := runCommand("apt-get", "install", "-y", "dnsmasq"); err != nil {
		fmt.Println("Error installing dnsmasq:", err)
		return
	}

	fmt.Println("dnsmasq installed successfully.")
}

