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
	// Step 1: Create the rsync configuration file at /etc/rsyncd.conf
	rsyncdConf := `
lock file = /var/run/rsync.lock
log file = /var/log/rsyncd.log
pid file = /var/run/rsyncd.pid

[code]
    path = /var/rsync
    comment = The documents folder of Juan
    uid = root
    gid = root
    read only = no
    list = yes
    auth users = root
    secrets file = /etc/rsyncd.secrets
    hosts allow = *
    ignore errors = yes
`
	err := ioutil.WriteFile("/etc/rsyncd.conf", []byte(rsyncdConf), 0644)
	if err != nil {
		fmt.Println("Error writing /etc/rsyncd.conf:", err)
		return
	}
	fmt.Println("Created /etc/rsyncd.conf")

	// Step 2: Create the rsync secrets file at /etc/rsyncd.secrets
	rsyncdSecrets := "root:password\n"
	err = ioutil.WriteFile("/etc/rsyncd.secrets", []byte(rsyncdSecrets), 0600)
	if err != nil {
		fmt.Println("Error writing /etc/rsyncd.secrets:", err)
		return
	}
	fmt.Println("Created /etc/rsyncd.secrets")

	// Step 3: Create the /var/rsync directory
	fmt.Println("Creating /var/rsync directory...")
	if err := os.MkdirAll("/var/rsync", 0755); err != nil {
		fmt.Println("Error creating /var/rsync directory:", err)
		return
	}

	// Step 4: Set permissions on /var/rsync directory
	fmt.Println("Setting permissions on /var/rsync directory...")
	if err := os.Chmod("/var/rsync", 0600); err != nil {
		fmt.Println("Error setting permissions on /var/rsync directory:", err)
		return
	}

	// Step 5: Start the rsync service
	fmt.Println("Starting rsync service...")
	if err := runCommand("service", "rsync", "start"); err != nil {
		fmt.Println("Error starting rsync service:", err)
		return
	}

	fmt.Println("Rsync setup completed successfully.")
}

