package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Define the host
	host := "google.com"

	// Run the ping command with -c 1 to send a single ping
	cmd := exec.Command("ping", "-c", "1", host)
	err := cmd.Run()

	// Check the return code (error)
	if err != nil {
		fmt.Printf("%s unreachable.\n", host)
	} else {
		fmt.Printf("%s reachable.\n", host)
	}
}

