package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ip := "8.8.8.8"
	timeout := 1 * time.Second

	// Attempt to resolve the address
	_, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		fmt.Println("Error resolving IP address:", err)
		return
	}

	// Attempt to establish a connection
	conn, err := net.DialTimeout("ip:icmp", ip, timeout)
	if err != nil {
		fmt.Println("IPv4 is down")
		return
	}
	defer conn.Close()

	fmt.Println("IPv4 is up")
}

