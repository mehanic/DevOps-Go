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

	// Define directories and files
	webDir := filepath.Join("/www/docs", domainName)
	confDir := "/etc/httpd/conf.d"
	confFile := filepath.Join(confDir, fmt.Sprintf("%s.conf", domainName))
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
	if err := ioutil.WriteFile(confFile, []byte(configContent), 0644); err != nil {
		fmt.Printf("Error writing config file: %s\n", err)
		os.Exit(1)
	}

	// Create the new website directory
	if err := os.MkdirAll(webDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating web directory: %s\n", err)
		os.Exit(1)
	}

	// Create the index.html file
	indexFilePath := filepath.Join(webDir, "index.html")
	if err := ioutil.WriteFile(indexFilePath, []byte(fmt.Sprintf("New site for %s", domainName)), 0644); err != nil {
		fmt.Printf("Error writing index.html: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("New site for %s created successfully.\n", domainName)

	// Ask if the user wants to restrict access
	var restrictAccess string
	fmt.Print("Do you want to restrict access to this site? (y/n): ")
	fmt.Scanln(&restrictAccess)

	if strings.ToLower(restrictAccess) == "n" {
		return
	}

	// Ask for the network to restrict access
	var network string
	fmt.Print("Which network should we restrict access to: ")
	fmt.Scanln(&network)

	// Append access restrictions to the configuration file
	restriction := fmt.Sprintf(`<Directory %s>
  Order allow,deny
  Allow from 127.0.0.1
  Allow from %s
</Directory>
`, webDir, network)

	// Read the existing config content
	currentConfig, err := ioutil.ReadFile(confFile)
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
		os.Exit(1)
	}

	// Find the closing tag and insert the restriction above it
	configContentWithRestriction := strings.Replace(string(currentConfig), "</VirtualHost>", restriction+"</VirtualHost>", 1)

	// Write the updated config content back to the file
	if err := ioutil.WriteFile(confFile, []byte(configContentWithRestriction), 0644); err != nil {
		fmt.Printf("Error updating config file: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Access restrictions added successfully.")
}

