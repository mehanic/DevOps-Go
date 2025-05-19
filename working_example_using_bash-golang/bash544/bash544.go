package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func dnsAxfr(domain string) {
	// Get the nameservers for the domain
	cmd := exec.Command("dig", domain, "ns", "+short")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error fetching nameservers: %s\n", err)
		return
	}

	nameservers := strings.Fields(string(output))

	for _, nameserver := range nameservers {
		// Perform zone transfer for each nameserver
		cmd = exec.Command("host", "-l", domain, nameserver)
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error fetching zone transfer for %s: %s\n", nameserver, err)
			continue
		}
		fmt.Print(string(output))
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[*] Simple Zone Transfer script")
		fmt.Println("[*] Usage: go run main.go <domain name>")
		os.Exit(0)
	}

	domain := os.Args[1]
	dnsAxfr(domain)
}

