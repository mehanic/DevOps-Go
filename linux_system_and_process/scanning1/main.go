package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

// Function to scan a single port
func socketScan(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this Goroutine as done once the function returns

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		fmt.Printf("[-] %d/tcp closed: %s\n", port, err.Error())
		return
	}
	fmt.Printf("[+] %d/tcp open\n", port)
	conn.Close()
}

// Function to scan multiple ports
func portScanning(host string, ports []int) {
	ipAddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		fmt.Printf("[-] Cannot resolve '%s': Unknown host\n", host)
		return
	}

	fmt.Printf("[+] Scan Results for: %s\n", ipAddr.String())

	var wg sync.WaitGroup
	for _, port := range ports {
		wg.Add(1)
		go socketScan(ipAddr.String(), port, &wg)
	}
	wg.Wait() // Wait for all Goroutines to finish
}

func main() {
	// Parsing command-line arguments
	host := flag.String("host", "localhost", "Host to scan")
	portRange := flag.String("ports", "80,443,8080", "Comma-separated list of ports to scan")
	flag.Parse()

	// Converting port list from string to integers
	var ports []int
	for _, p := range strings.Split(*portRange, ",") {
		var port int
		fmt.Sscanf(p, "%d", &port)
		ports = append(ports, port)
	}

	// Start scanning
	portScanning(*host, ports)
}
