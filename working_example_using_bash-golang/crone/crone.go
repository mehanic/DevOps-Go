package main

import (
	"fmt"
	"os"
	"time"
)

// logDate appends the current date and time to the specified file.
func logDate(filePath string) error {
	// Open the file in append mode, create it if it doesn't exist.
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Get the current date and time.
	currentTime := time.Now().Format(time.RFC3339) // Format: YYYY-MM-DD HH:MM:SS
	// Write the date and time to the file.
	if _, err := file.WriteString(currentTime + "\n"); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

func main() {
	const filePath = "/tmp/date-file"

	// Create a ticker that ticks every minute.
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	fmt.Println("Starting the date logger...")

	for {
		select {
		case <-ticker.C:
			if err := logDate(filePath); err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}

