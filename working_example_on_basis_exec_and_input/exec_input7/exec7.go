package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() { // Open the configuration file for editing
	cmd := exec.Command("nano", "/etc/network/interfaces")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	// Restart the network service to apply the changes
	cmd = exec.Command("service", "networking", "restart")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Base configurations modified successfully!")
}
