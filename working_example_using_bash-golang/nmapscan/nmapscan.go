package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run script.go subnet")
		os.Exit(1)
	}

	subnet := os.Args[1]
	opts := []string{"-T4", "--open"}
	pingopts := []string{"-sn", "-PS21-23,25,53,80,443,3389", "-PO", "-PE", "-PM", "-PP"}

	// Capture SIGINT and SIGTERM to exit gracefully
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigs)

	fmt.Println("--------")
	fmt.Println("Finding active hosts")
	fmt.Println("--------")

	// Construct the nmap command for host discovery
	cmd := exec.Command("nmap", append(opts, append(pingopts, "-oG", "alive.gnmap", subnet)...)...)
	fmt.Printf("Running command: %s\n", strings.Join(cmd.Args, " "))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing nmap: %v\nOutput: %s\n", err, output)
		os.Exit(1)
	}

	// Read alive.gnmap to find active hosts
	file, err := os.Open("alive.gnmap")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening alive.gnmap: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var targets []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Status: Up") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				targets = append(targets, fields[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading alive.gnmap: %v\n", err)
		os.Exit(1)
	}

	count := len(targets)
	fmt.Printf("[+] Found %d active hosts.\n\n", count)

	// Check if any active hosts were found
	if count == 0 {
		fmt.Println("No active hosts found. Exiting.")
		return
	}

	fmt.Println("--------")
	fmt.Println("Finding open ports")
	fmt.Println("--------")

	// Create targets.txt file for nmap input
	targetsFile, err := os.Create("targets")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating targets file: %v\n", err)
		os.Exit(1)
	}
	defer targetsFile.Close()

	for _, target := range targets {
		targetsFile.WriteString(target + "\n")
	}

	// Construct the nmap command for open port scanning
	cmd = exec.Command("nmap", append(opts, "-iL", "targets", "-p", "1-65535", "-oG", "ports.gnmap")...)
	fmt.Printf("Running command: %s\n", strings.Join(cmd.Args, " "))
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing nmap for ports: %v\nOutput: %s\n", err, output)
		os.Exit(1)
	}

	// Read ports.gnmap to find open ports
	file, err = os.Open("ports.gnmap")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening ports.gnmap: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	openPorts := make(map[string]struct{})
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "/open") {
			port := strings.Split(line, "/")[0]
			openPorts[port] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading ports.gnmap: %v\n", err)
		os.Exit(1)
	}

	count = len(openPorts)
	fmt.Printf("[+] Found %d unique open ports\n\n", count)

	// Check if any open ports were found
	if count == 0 {
		fmt.Println("No open ports found. Exiting.")
		return
	}

	// Prepare the port list for the full scan
	var portList []string
	for port := range openPorts {
		portList = append(portList, port)
	}

	// Construct the nmap command for full scan
	fmt.Println("--------")
	fmt.Println("Running full nmap scan")
	fmt.Println("--------")

	portListStr := strings.Join(portList, ",")
	cmd = exec.Command("nmap", append(opts, "-iL", "targets", "-p", portListStr, "-A", "-oA", "full_scan")...)
	fmt.Printf("Running command: %s\n", strings.Join(cmd.Args, " "))
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing full nmap scan: %v\nOutput: %s\n", err, output)
		os.Exit(1)
	}

	fmt.Println("[+] Scan results available in full_scan.*")
}

