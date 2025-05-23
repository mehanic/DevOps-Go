package main

import (
	"fmt"
	"os"
	"os/exec"
//	"path/filepath"
)

// Utility function to run a command with arguments
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Step 1: Install NFS server and common packages
	fmt.Println("Installing NFS server and common packages...")
	if err := runCommand("apt-get", "install", "-y", "nfs-kernel-server", "nfs-common"); err != nil {
		fmt.Println("Error installing NFS packages:", err)
		return
	}

	// Step 2: Create the /var/nfs directory
	nfsDir := "/var/nfs"
	fmt.Println("Creating NFS directory:", nfsDir)
	if err := os.MkdirAll(nfsDir, 0755); err != nil {
		fmt.Println("Error creating NFS directory:", err)
		return
	}

	// Step 3: Append export configuration to /etc/exports
	exportsFile := "/etc/exports"
	exportLine := "/var/nfs * (rw,no_root_squash)\n"
	fmt.Println("Updating", exportsFile, "with NFS export entry.")
	file, err := os.OpenFile(exportsFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening /etc/exports:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(exportLine); err != nil {
		fmt.Println("Error writing to /etc/exports:", err)
		return
	}

	// Step 4: Start rpcbind service
	fmt.Println("Starting rpcbind service...")
	if err := runCommand("systemctl", "start", "rpcbind"); err != nil {
		fmt.Println("Error starting rpcbind service:", err)
		return
	}

	// Step 5: Start nfs-server service
	fmt.Println("Starting NFS server service...")
	if err := runCommand("systemctl", "start", "nfs-server"); err != nil {
		fmt.Println("Error starting NFS server service:", err)
		return
	}

	fmt.Println("NFS server setup completed successfully.")
}

