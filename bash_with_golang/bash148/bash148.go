package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
)

// logMessage logs a message to the system log
func logMessage(message string, priority syslog.Priority, tag string, includeStdErr bool) {
	var logger *syslog.Writer
	var err error

	// Connect to the syslog
	if tag != "" {
		logger, err = syslog.New(priority, tag)
	} else {
		logger, err = syslog.New(priority, "")
	}

	if err != nil {
		log.Fatalf("Failed to connect to syslog: %v", err)
	}
	defer logger.Close()

	// Log the message
	if includeStdErr {
		// Print to stderr in addition to syslog
		fmt.Fprintln(os.Stderr, message)
	}

	err = logger.Info(message)
	if err != nil {
		log.Fatalf("Failed to log message: %v", err)
	}
}

func main() {
	// Mimicking the various logger commands

	// logger "Message"
	logMessage("Message", syslog.LOG_INFO|syslog.LOG_LOCAL0, "", false)

	// logger -p local0.info "Message"
	logMessage("Message", syslog.LOG_INFO|syslog.LOG_LOCAL0, "", false)

	// logger -s -p local0.info "Message"
	logMessage("Message", syslog.LOG_INFO|syslog.LOG_LOCAL0, "", true)

	// logger -t myscript -p local0.info "Message"
	logMessage("Message", syslog.LOG_INFO|syslog.LOG_LOCAL0, "myscript", false)

	// logger -i -t myscript "Message"
	logMessage("Message", syslog.LOG_INFO|syslog.LOG_LOCAL0, "myscript", false)
}

