package main

import (
	"fmt"
//	"net"
	"os"
	"strings"
)

func isValidIP(ipAddr string) bool {
	// Split the IP address into octets
	octets := strings.Split(ipAddr, ".")
	if len(octets) != 4 {
		return false
	}

	// Validate each octet
	for _, octet := range octets {
		if !isValidOctet(octet) {
			return false
		}
	}

	return true
}

func isValidOctet(octet string) bool {
	// Check if octet is a valid integer
	var num int
	n, err := fmt.Sscanf(octet, "%d", &num)
	if n != 1 || err != nil {
		return false
	}
	// Check if octet is in the valid range
	return num >= 0 && num <= 255
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go <IP_ADDRESS>")
		return
	}

	ipAddr := os.Args[1]

	if isValidIP(ipAddr) {
		fmt.Printf("%s is valid\n", ipAddr)
	} else {
		fmt.Printf("%s is NOT valid\n", ipAddr)
	}
}

