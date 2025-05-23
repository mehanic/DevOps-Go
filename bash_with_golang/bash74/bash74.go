package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Execute the ping command with -q (quiet), -c 1 (1 ping), and -W 1 (timeout 1 second)
	cmd := exec.Command("ping", "-q", "-c", "1", "-W", "1", "8.8.8.8")

	// Run the command and check for errors
	err := cmd.Run()

	// If no error, the ping was successful, otherwise it failed
	if err == nil {
		fmt.Println("IPv4 is up")
	} else {
		fmt.Println("IPv4 is down")
	}
}

