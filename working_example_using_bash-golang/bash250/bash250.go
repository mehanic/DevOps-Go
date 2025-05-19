package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Check if a domain name is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <domain_name>")
		os.Exit(1)
	}

	domainName := os.Args[1]

	// Define directories
	webDir := "/www/docs"
	confDir := "/etc/httpd/conf.d"
	template := filepath.Join(os.Getenv("HOME"), "template.txt")

	// Ensure the configuration directory exists
	if err := os.MkdirAll(confDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating config directory: %s\n", err)
		os.Exit(1)
	}

	// Read the template file
	templateContent, err := ioutil.ReadFile(template)
	if err != nil {
		fmt.Printf("Error reading template file: %s\n", err)
		os.Exit(1)
	}

	// Replace the placeholder with the provided domain name
	configContent := strings.ReplaceAll(string(templateContent), "dummy-host.example.com", domainName)

	// Write the new configuration file
	confFilePath := filepath.Join(confDir, fmt.Sprintf("%s.conf", domainName))
	if err := ioutil.WriteFile(confFilePath, []byte(configContent), 0644); err != nil {
		fmt.Printf("Error writing config file: %s\n", err)
		os.Exit(1)
	}

	// Create the new website directory
	siteDir := filepath.Join(webDir, domainName)
	if err := os.MkdirAll(siteDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating web directory: %s\n", err)
		os.Exit(1)
	}

	// Create the index.html file
	indexFilePath := filepath.Join(siteDir, "index.html")
	if err := ioutil.WriteFile(indexFilePath, []byte(fmt.Sprintf("New site for %s", domainName)), 0644); err != nil {
		fmt.Printf("Error writing index.html: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("New site for %s created successfully.\n", domainName)
}

