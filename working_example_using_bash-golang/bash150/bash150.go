package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
//	"os/exec"
	"time"
)

var (
	VERBOSE    = false           // Verbosity flag
	HOST       = "https://www.google.com" // The host to fetch data from
	PID        = os.Getpid()     // Process ID of the current program
	PROGRAM_NAME = os.Args[0]    // Name of the current program
	THIS_HOST  string            // The hostname of the machine
)

// logit logs messages with a timestamp, host, program name, and PID
func logit(logLevel string, msg string) {
	TIMESTAMP := time.Now().Format("2006-01-02 15:04:05")
	if logLevel == "ERROR" || VERBOSE {
		fmt.Printf("%s %s %s[%d]: %s %s\n", TIMESTAMP, THIS_HOST, PROGRAM_NAME, PID, logLevel, msg)
	}
}

func fetchData(host string) error {
	resp, err := http.Get(host) // Fetching data from the host
	if err != nil {
		return err // Return error if fetching fails
	}
	defer resp.Body.Close() // Ensure the response body is closed

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data: %s", resp.Status) // Check for successful response
	}

	return nil // Return nil if fetching was successful
}

func main() {
	// Get the hostname of the machine
	var err error
	THIS_HOST, err = os.Hostname()
	if err != nil {
		log.Fatalf("Could not get hostname: %v", err)
	}

	logit("INFO", "Processing data.")

	// Try to fetch data from the host
	if err := fetchData(HOST); err != nil {
		logit("ERROR", fmt.Sprintf("Could not fetch data from %s: %v", HOST, err))
	}
}

