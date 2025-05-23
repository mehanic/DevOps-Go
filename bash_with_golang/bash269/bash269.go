package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open /etc/hosts file for reading
	file, err := os.Open("/etc/hosts")
	if err != nil {
		fmt.Println("Error opening /etc/hosts:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read the current line
		line := scanner.Text()

		// Skip commented or empty lines
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		// Split the line into fields (IP, name, and possible aliases)
		fields := strings.Fields(line)
		if len(fields) < 2 {
			// Skip lines that don't have at least IP and name
			continue
		}

		ip := fields[0]
		name := fields[1]

		// Join the remaining fields as aliases (if any exist)
		aliases := strings.Join(fields[2:], " ")

		// Print the output in the desired format
		fmt.Printf("IP is %s - its name is %s", ip, name)
		if aliases != "" {
			fmt.Printf("   Aliases: %s\n", aliases)
		} else {
			fmt.Println() // Just print a blank line if no aliases
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}

