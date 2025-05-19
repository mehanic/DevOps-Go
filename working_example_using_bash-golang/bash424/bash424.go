package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	hostname := "google.com"
	timeout := 1 * time.Second

	// Attempt to resolve the hostname
	_, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println("The network is down")
		return
	}

	// Attempt to establish a connection
	conn, err := net.DialTimeout("tcp", hostname+":80", timeout) // Check connectivity to port 80 (HTTP)
	if err != nil {
		fmt.Println("The network is down")
		return
	}
	defer conn.Close()

	fmt.Println("The network is up")
}

