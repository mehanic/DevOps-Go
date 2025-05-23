package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	urlFile      = "Host_PortFile.txt"
	port         = "80" // Default port
	mailTo       = "admin@techpaste.com" // Replace (at) and (dot) with @ and .
	logFileName  = "script.log"
	pingAttempts  = 6
	pingInterval  = 2 * time.Second
	telnetTimeout = 5 * time.Second
)

func main() {
	// Set parameters
	timeStamp := time.Now().Format("02-01-2006_15.04.05")

	// Create or open the log file
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Read the URL and port file
	file, err := os.Open(urlFile)
	if err != nil {
		fmt.Println("Error opening URL file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Ping hosts
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ":")
		server := parts[0]

		if pingHost(server) {
			logStatus(logFile, timeStamp, server, "[ RUNNING ]")
		} else {
			logStatus(logFile, timeStamp, server, "[ DOWN ]")
			sendEmail(fmt.Sprintf("%s : Status Of Host %s = [ DOWN ]", timeStamp, server))
		}
	}

	// Reset scanner for next operation
	file.Seek(0, 0)

	// Check Telnet status
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.Split(line, ":")
		server := parts[0]
		var port string
		if len(parts) > 1 {
			port = parts[1]
		} else {
			port = "80" // Default port
		}

		if checkTelnet(server, port) {
			logStatus(logFile, timeStamp, fmt.Sprintf("http://%s:%s/", server, port), "[ OPEN ]")
		} else {
			logStatus(logFile, timeStamp, fmt.Sprintf("http://%s:%s/", server, port), "[ NOT OPEN ]")
			sendEmail(fmt.Sprintf("%s : Port %s of URL %s is NOT OPEN", timeStamp, port, fmt.Sprintf("http://%s:%s/", server, port)))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}
}

// pingHost pings the given server and returns true if it is reachable.
func pingHost(server string) bool {
	cmd := exec.Command("ping", "-c", fmt.Sprint(pingAttempts), "-i", fmt.Sprint(pingInterval.Seconds()), server)
	err := cmd.Run()
	return err == nil
}

// checkTelnet checks if the specified port on the server is open using a TCP connection.
func checkTelnet(server, port string) bool {
	address := fmt.Sprintf("%s:%s", server, port)
	conn, err := net.DialTimeout("tcp", address, telnetTimeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// logStatus logs the status of the host/port to the log file.
func logStatus(logFile *os.File, timeStamp, identifier, status string) {
	logEntry := fmt.Sprintf("%s : %s = %s\n", timeStamp, identifier, status)
	fmt.Print(logEntry)
	if _, err := logFile.WriteString(logEntry); err != nil {
		fmt.Println("Error writing to log file:", err)
	}
}

// sendEmail sends an email notification (dummy implementation).
func sendEmail(body string) {
	// This function would normally send an email, but here we will just print it.
	// You can use an SMTP library or command-line email utility to implement this.
	fmt.Printf("Sending email to %s with body: %s\n", mailTo, body)
}

