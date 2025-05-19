package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func getUptime() (string, error) {
	// Run the uptime command
	cmd := exec.Command("uptime")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Convert output to string
	uptimeStr := string(output)

	// Check for days in the uptime string
	if strings.Contains(uptimeStr, "day") {
		// Extract the days and hours
		parts := strings.Fields(uptimeStr)
		if len(parts) >= 4 {
			return strings.Join(parts[2:4], " "), nil
		}
	} else {
		// Extract just the uptime (without days)
		parts := strings.Fields(uptimeStr)
		if len(parts) >= 3 {
			return strings.Trim(parts[2], ","), nil
		}
	}

	return "", nil
}

func main() {
	serverUptime, err := getUptime()
	if err != nil {
		fmt.Println("Error retrieving uptime:", err)
		return
	}

	fmt.Println("Server Uptime:", serverUptime)
}

