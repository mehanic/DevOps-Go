package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	for {
		// Execute the ping command
		cmd := exec.Command("ping", "-c", "1", "app1")
		err := cmd.Run()

		if err != nil {
			// If the ping command fails, app1 is down
			break
		}

		// If ping is successful, app1 is still up
		fmt.Println("app1 still up...")
		time.Sleep(5 * time.Second) // Sleep for 5 seconds
	}

	fmt.Println("app1 down, continuing.")
}

