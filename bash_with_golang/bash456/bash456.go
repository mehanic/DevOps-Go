package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Initialize the config map
	config := map[string]string{
		"username": "student",
		"password": "",
		"hostname": "ubuntu",
	}

	// Open the configuration file
	file, err := os.Open("sampleconfig.conf")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			varname := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			config[varname] = value // Set the value in the config map
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Print the configuration values
	fmt.Println(config["username"])
	fmt.Println(config["password"])
	fmt.Println(config["hostname"])
	fmt.Println(config["PROMPT_COMMAND"]) // Will be empty if not set in the file
}

