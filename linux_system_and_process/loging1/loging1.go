package main

import (
//	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Custom logger with different log levels
func logWithLevel(level, msg string) {
	// Get the file and line number
	_, file, line, _ := runtime.Caller(2)
	// Get the timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// Log message with timestamp, level, file, and line number
	log.Printf("%s - [%s] (%s:%d) %s\n", timestamp, level, file, line, msg)
}

func main() {
	// Configure logging to print to a file (optional, here we use stdout)
	log.SetOutput(os.Stdout)

	// Simulating different log levels
	logWithLevel("DEBUG", "Some additional information")
	logWithLevel("INFO", "Working...")
	logWithLevel("WARNING", "Watch out!")
	logWithLevel("ERROR", "Oh NO!")
	logWithLevel("CRITICAL", "x.x")
}

// 2025/01/19 13:52:44 2025-01-19 13:52:44 - [DEBUG] (/home/mehanic/.gvm/gos/go1.23.0/src/runtime/proc.go:272) Some additional information
// 2025/01/19 13:52:44 2025-01-19 13:52:44 - [INFO] (/home/mehanic/.gvm/gos/go1.23.0/src/runtime/proc.go:272) Working...
// 2025/01/19 13:52:44 2025-01-19 13:52:44 - [WARNING] (/home/mehanic/.gvm/gos/go1.23.0/src/runtime/proc.go:272) Watch out!
// 2025/01/19 13:52:44 2025-01-19 13:52:44 - [ERROR] (/home/mehanic/.gvm/gos/go1.23.0/src/runtime/proc.go:272) Oh NO!
// 2025/01/19 13:52:44 2025-01-19 13:52:44 - [CRITICAL] (/home/mehanic/.gvm/gos/go1.23.0/src/runtime/proc.go:272) x.x
