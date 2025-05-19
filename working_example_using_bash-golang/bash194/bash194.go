package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/go-vgo/robotgo"
)

func main() {
    fmt.Println("Press Ctrl-C to quit.")

    // Create a channel to listen for interrupt signals
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        for {
            // Get mouse coordinates
            x, y := robotgo.GetMousePos()

            // Get the pixel color at the mouse position
            pixelColor := robotgo.GetPixelColor(x, y)

            // Format the output string
            positionStr := fmt.Sprintf("X: %4d Y: %4d RGB: %s", x, y, pixelColor)

            // Clear the line by moving the cursor back
            fmt.Print(positionStr)
            fmt.Print("\r") // Move the cursor back to the beginning of the line

            // Wait for a short duration before updating the position again
            time.Sleep(100 * time.Millisecond)
        }
    }()

    // Wait for the interrupt signal
    <-sigChan
    fmt.Println("\nDone.")
}

