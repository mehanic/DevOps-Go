package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func getCPUInfo() (string, error) {
	// Run the dmidecode command
	cmd := exec.Command("dmidecode", "-t4")
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command failed: %s, error: %v", stderr.String(), err)
	}

	return out.String(), nil
}

func parseCPUInfo(cpuInfo string) {
	// Split the output into lines for easier processing
	lines := strings.Split(cpuInfo, "\n")

	// Use a regex to extract relevant information
	var cpuHandles []string
	for _, line := range lines {
		if strings.HasPrefix(line, "Handle") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				cpuHandles = append(cpuHandles, parts[1])
			}
		}
	}

	// Loop through each CPU handle and retrieve information
	for _, handle := range cpuHandles {
		fmt.Println("Handle:", handle)

		// Extract details using a regex pattern
		cpuDetailPattern := regexp.MustCompile(fmt.Sprintf("Handle %s(.*?)(?=Handle|$)", handle))
		details := cpuDetailPattern.FindStringSubmatch(cpuInfo)

		if len(details) > 0 {
			// Get the relevant section for this handle
			cpuDetails := details[1]

			// Define the labels we're interested in
			labels := []string{"Socket Designation", "Family", "Manufacturer", "Current Speed", "Voltage", "Core Count"}

			for _, label := range labels {
				// Search for each label and retrieve its value
				re := regexp.MustCompile(fmt.Sprintf("%s: (.*?)\n", label))
				matches := re.FindStringSubmatch(cpuDetails)
				if len(matches) > 1 {
					// Print the label and its value
					fmt.Printf("%s: %s\n", label, strings.TrimSpace(matches[1]))
				}
			}
		}
		fmt.Println() // Add an empty line between CPUs
	}
}

func main() {
	cpuInfo, err := getCPUInfo()
	if err != nil {
		if strings.Contains(err.Error(), "permission denied") {
			fmt.Println("Please run this program as root or using sudo.")
		} else {
			fmt.Println("Error retrieving CPU information:", err)
		}
		return
	}

	parseCPUInfo(cpuInfo)
}

