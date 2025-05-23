package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
//	"strings"
)

func main() {
	// Check if the correct number of arguments is provided
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go filename")
		os.Exit(1)
	}

	// Read the filename from command line arguments
	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read hosts from the file
	var hosts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hosts = append(hosts, scanner.Text())
	}

	// Check for errors during file reading
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	gb := "gobuster" // Gobuster command
	dl := "/usr/share/seclists/Discovery/Web_Content/raft-small-directories.txt"
	wl := "/usr/share/seclists/Discovery/Web_Content/raft-small-files.txt"

	// Run Gobuster against each URL in the hosts list
	for _, host := range hosts {
		fmt.Printf("Busting %s\n", host)

		// Run Gobuster with the directory wordlist
		cmd1 := exec.Command(gb, "-q", "-e", "-s", "200", "-m", "dir", "-u", host, "-w", dl)
		output1, err := cmd1.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running gobuster for %s: %v\n", host, err)
		}
		fmt.Printf("%s\n", output1)

		// Run Gobuster with the file wordlist
		cmd2 := exec.Command(gb, "-q", "-e", "-s", "200", "-m", "dir", "-u", host, "-w", wl)
		output2, err := cmd2.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running gobuster for %s: %v\n", host, err)
		}
		fmt.Printf("%s\n", output2)

		fmt.Println("")
	}
}

