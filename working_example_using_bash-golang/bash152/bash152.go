package main

import (
	"fmt"
	"log"
	"log/syslog"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Call my_logger with random numbers
	myLogger(fmt.Sprintf("Random number: %d", rand.Intn(10000)))
	myLogger(fmt.Sprintf("Random number: %d", rand.Intn(10000)))
	myLogger(fmt.Sprintf("Random number: %d", rand.Intn(10000)))
}

// myLogger logs a message to the console and to syslog
func myLogger(message string) {
	// Print the message to the console
	fmt.Println(message)

	// Log the message to syslog
	logMessage(message)
}

// logMessage logs a message to the system log with a specified tag
func logMessage(message string) {
	// Connect to syslog with a specified tag and priority
	logger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_USER, "randomly")
	if err != nil {
		log.Fatalf("Failed to connect to syslog: %v", err)
	}
	defer logger.Close() // Ensure the logger is closed when done

	// Log the message
	err = logger.Info(message)
	if err != nil {
		log.Fatalf("Failed to log message: %v", err)
	}
}

