package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Press Ctrl-C to quit.")

	// Create a ticker to update every 100 milliseconds (0.1 seconds)
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	// Run indefinitely until interrupted
	for {
		select {
		case <-ticker.C:
			// Get the current mouse position
			x, y := robotgo.GetMousePos()

			// Get the pixel color at the mouse position
			color := robotgo.GetPixelColor(x, y)

			// Format and display the mouse position and pixel color
			positionStr := fmt.Sprintf("X: %4d Y: %4d RGB: (%3s)", x, y, color)
			fmt.Print(positionStr)
			// Move the cursor back to the beginning of the line
			fmt.Print("\r")

		}
	}
}

