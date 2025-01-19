package main

import (
	"fmt"
	"os"
	"log"
	"regexp"
)

// Define a function to extract IP addresses from a log file
func extractIPsFromLog(logFilePath string) []string {
	// Read the log file's content
	data, err := os.ReadFile(logFilePath)
	if err != nil {
		log.Fatalf("Failed to read the log file: %s", err)
	}

	// Regular expression to match IP addresses
	re := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)

	// Find all matches in the log data
	ipAddresses := re.FindAllString(string(data), -1)

	return ipAddresses
}

func main() {
	// Call the function and print the extracted IP addresses
	ips := extractIPsFromLog("sample_log.txt")
	fmt.Println("Extracted IP Addresses:")
	for _, ip := range ips {
		fmt.Println(ip)
	}
}


// go run main.go < sample_log.txt                                                                                                                              
// Extracted IP Addresses:
// 192.168.1.1
// 10.0.0.2
// 172.16.0.3
