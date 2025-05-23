package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Check number of arguments passed
	if len(os.Args) < 3 {
		fmt.Println("Error, usage is:")
		fmt.Println("ftpget hostname filename [directory].")
		os.Exit(-1)
	}

	hostname := os.Args[1]
	filename := os.Args[2]
	directory := "." // Default value

	if len(os.Args) >= 4 {
		directory = os.Args[3]
	}

	// Create the FTP command sequence
	ftpCommands := fmt.Sprintf("open %s\ncd %s\nget %s\nquit\n", hostname, directory, filename)

	// Execute the FTP command
	cmd := exec.Command("ftp")

	// Provide the FTP commands via standard input
	cmd.Stdin = strings.NewReader(ftpCommands)

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Println("Error starting FTP session:", err)
		return
	}

	fmt.Println("FTP session ended.")
}

