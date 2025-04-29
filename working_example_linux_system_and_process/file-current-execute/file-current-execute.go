package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    // Retrieve the path of the currently running executable
    path, err := os.Executable()
    if err != nil {
        // Log the error if the executable path could not be retrieved
        log.Println("Error retrieving executable path:", err)
        return
    }

    // Print the path of the executable
    fmt.Println("Executable path:", path)

    // If an error occurred, log the error and exit the program
    if err != nil {
        log.Fatal("Exiting due to error:", err)
    }
}