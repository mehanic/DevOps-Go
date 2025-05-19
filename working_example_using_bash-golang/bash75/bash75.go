package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Execute the ping command with -q (quiet), -c 1 (1 ping), and -W 1 (timeout 1 second)
	cmd := exec.Command("ping", "-q", "-c", "1", "-W", "1", "google.com")

	// Run the command and check for errors
	err := cmd.Run()

	// If no error, the network is up, otherwise it's down
	if err == nil {
		fmt.Println("The network is up")
	} else {
		fmt.Println("The network is down")
	}
}

