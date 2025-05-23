package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		// Print message
		fmt.Println("Shall run for ever")

		// Sleep for 1 second (similar to `sleep 1` in Bash)
		time.Sleep(1 * time.Second)
	}
}

