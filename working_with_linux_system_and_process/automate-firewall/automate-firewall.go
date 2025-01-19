package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Open port 80 for HTTP traffic
	cmd1 := exec.Command("iptables", "-A", "INPUT", "-p",
		"tcp", "--dport", "80", "-j", "ACCEPT")
	err := cmd1.Run()
	if err != nil {
		fmt.Println("Error opening port 80:", err)
		return
	}
	// Open port 443 for HTTPS traffic
	cmd2 := exec.Command("iptables", "-A", "INPUT", "-p",
		"tcp", "--dport", "443", "-j", "ACCEPT")
	err = cmd2.Run()
	if err != nil {
		fmt.Println("Error opening port 443:", err)
		return
	}
	// Deny all other incoming traffic
	cmd3 := exec.Command("iptables", "-A", "INPUT", "-j",
		"DROP")
	err = cmd3.Run()
	if err != nil {
		fmt.Println("Error blocking incoming traffic:", err)
		return
	}
	// Display the current iptables rules
	cmd4 := exec.Command("iptables", "-L")
	output, err := cmd4.Output()
	if err != nil {
		fmt.Println("Error getting iptables rules:", err)
		return
	}
	fmt.Println("Current iptables rules:\n", string(output))
}


//Error opening port 80: exit status 4

