package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// Run the uptime command
	cmd := exec.Command("uptime")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing uptime command:", err)
		return
	}

	// Convert output to string and split it by spaces
	uptimeOutput := string(output)
	fields := strings.Fields(uptimeOutput)

	// Check if the output contains "day"
	var serverUptime string
	if len(fields) >= 4 && strings.Contains(fields[3], "day") {
		// If "day" is in the output, concatenate fields 3 and 4
		serverUptime = fields[2] + " " + strings.Trim(fields[3], ",")
	} else {
		// Otherwise, use only field 3
		serverUptime = strings.Trim(fields[2], ",")
	}

	// Print the server uptime
	fmt.Println(serverUptime)
}

