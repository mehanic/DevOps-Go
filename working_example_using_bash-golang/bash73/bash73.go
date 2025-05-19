package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func runDmidecode(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func getCpuHandles() ([]string, error) {
	// Get all CPU handles using dmidecode
	output, err := runDmidecode("dmidecode -t4 | awk '/Handle / {print $2}'")
	if err != nil {
		return nil, err
	}
	// Split the output by newline and return as a slice of handles
	handles := strings.Split(strings.TrimSpace(output), "\n")
	return handles, nil
}

func extractField(handle, field string) (string, error) {
	// Run dmidecode and extract relevant fields using the handle
	command := fmt.Sprintf("dmidecode -t4 | sed '/Flags/,/Version/d' | egrep -A20 'Handle %s' | grep -m 1 '%s'", handle, field)
	output, err := runDmidecode(command)
	if err != nil {
		return "", err
	}
	
	// Use regex to extract the value after the colon
	re := regexp.MustCompile(`: *(.{0,18})`)
	match := re.FindStringSubmatch(output)
	if len(match) > 1 {
		return strings.TrimSpace(match[1]), nil
	}
	return "", nil
}

func main() {
	// Get CPU handles
	handles, err := getCpuHandles()
	if err != nil {
		fmt.Println("Error fetching CPU handles:", err)
		return
	}

	// Loop through each CPU handle and extract details
	for _, handle := range handles {
		fmt.Printf("Handle: %s\n", handle)

		socketDesignation, err := extractField(handle, "Socket Designation")
		if err == nil {
			fmt.Println("Socket Designation:", socketDesignation)
		}

		family, err := extractField(handle, "Family")
		if err == nil {
			fmt.Println("Family:", family)
		}

		manufacturer, err := extractField(handle, "Manufacturer")
		if err == nil {
			fmt.Println("Manufacturer:", manufacturer)
		}

		currentSpeed, err := extractField(handle, "Current Speed")
		if err == nil {
			fmt.Println("Current Speed:", currentSpeed)
		}

		voltage, err := extractField(handle, "Voltage")
		if err == nil {
			fmt.Println("Voltage:", voltage)
		}

		coreCount, err := extractField(handle, "Core Count")
		if err == nil {
			fmt.Println("Core Count:", coreCount)
		}
		
		fmt.Println()
	}
}

