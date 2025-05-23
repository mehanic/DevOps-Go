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
	randomNumber := rand.Intn(10000) // Generate a random number between 0 and 9999
	message := fmt.Sprintf("Random number: %d", randomNumber)

	// Print the message to the console
	fmt.Println(message)

	// Log the message to syslog
	logMessage(message)
}

// logMessage logs a message to the system log
func logMessage(message string) {
	// Connect to syslog
	logger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_USER, "mygoapp")
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

