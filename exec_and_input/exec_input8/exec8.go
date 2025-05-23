package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// // Update the package lists
	// cmd := exec.Command("sudo", "apt-get", "update", "-y", "--allow-unauthenticated update")
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println("Error updating package lists:", err)
	// 	return
	// }
	// Upgrade the packages
	cmd := exec.Command("sudo", "apt-get", "upgrade", "-y")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error upgrading packages:", err)
		return
	}
	fmt.Println("System is up-to-date.")
}

