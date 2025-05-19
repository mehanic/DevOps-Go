package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"net/http"
	"os"
	"strings"
)

// Replace with your Shodan API key
const apiKey = "YOUR_SHODAN_API_KEY"

// ShodanHost represents the structure of the response from Shodan API
type ShodanHost struct {
	OS        string   `json:"os"`
	Hostnames []string `json:"hostnames"`
	Ports     []int    `json:"ports"`
	Vulns     []string `json:"vulns"`
}

// lookupIP queries the Shodan API for the given IP and prints the results
func lookupIP(ip string) {
	fmt.Println(ip)
	fmt.Println(strings.Repeat("=", len(ip)))

	url := fmt.Sprintf("https://api.shodan.io/shodan/host/%s?key=%s", ip, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("API Error: %v\n\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n\n", err)
		return
	}

	// Parse JSON response
	var host ShodanHost
	err = json.Unmarshal(body, &host)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n\n", err)
		return
	}

	// Print results
	fmt.Printf("Operating System: %s\n", host.OS)
	fmt.Printf("Hostnames: %s\n", strings.Join(host.Hostnames, ", "))
	fmt.Printf("Ports: %s\n", strings.Join(intSliceToStringSlice(host.Ports), ", "))
	fmt.Printf("Vulns: %s\n\n", strings.Join(host.Vulns, ", "))
}

// intSliceToStringSlice converts a slice of ints to a slice of strings
func intSliceToStringSlice(ints []int) []string {
	var strs []string
	for _, i := range ints {
		strs = append(strs, fmt.Sprintf("%d", i))
	}
	return strs
}

// incIP increments the given IP address by one
func incIP(ip net.IP) net.IP {
	ip = ip.To16() // Ensure IPv6 or IPv4 as IPv6-mapped address
	intIP := big.NewInt(0)
	intIP.SetBytes(ip)

	// Increment the integer
	intIP.Add(intIP, big.NewInt(1))

	// Convert back to net.IP
	return net.IP(intIP.Bytes())
}

// getNetworkAddress returns the network address given an IPNet
func getNetworkAddress(ipnet *net.IPNet) net.IP {
	ip := ipnet.IP.Mask(ipnet.Mask)
	return ip
}

func main() {
	// Parse command-line arguments
	var singleIP string
	var startIP, endIP string
	var cidrRange string

	flag.StringVar(&singleIP, "i", "", "A single IP address. ex: 192.168.1.1")
	flag.StringVar(&startIP, "r-start", "", "Start of an IP range. ex: 192.168.1.1")
	flag.StringVar(&endIP, "r-end", "", "End of an IP range. ex: 192.168.1.10")
	flag.StringVar(&cidrRange, "c", "", "A range of IPs in CIDR notation. ex: 192.168.1.0/24")
	flag.Parse()

	// Determine the input method
	if singleIP != "" {
		lookupIP(singleIP)

	} else if startIP != "" && endIP != "" {
		// Handle IP range
		start := net.ParseIP(startIP)
		end := net.ParseIP(endIP)
		if start == nil || end == nil {
			fmt.Println("Invalid IP address format.")
			os.Exit(1)
		}

		// Get all IPs in range
		for ip := start; !ip.Equal(end); ip = incIP(ip) {
			lookupIP(ip.String())
		}
		// Lookup the last IP as well
		lookupIP(end.String())

	} else if cidrRange != "" {
		// Handle CIDR notation
		_, ipnet, err := net.ParseCIDR(cidrRange)
		if err != nil {
			fmt.Printf("Address Error: %v.\n", err)
			os.Exit(1)
		}

		// Iterate over the range of IPs in the CIDR block
		for ip := getNetworkAddress(ipnet); ipnet.Contains(ip); ip = incIP(ip) {
			lookupIP(ip.String())
		}

	} else {
		fmt.Println("No valid input provided.")
		flag.Usage()
		os.Exit(1)
	}
}

