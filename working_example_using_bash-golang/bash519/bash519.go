package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Clear the terminal window
	clearTerminal()

	// Greet the user
	logName := os.Getenv("LOGNAME")
	if logName == "" {
		logName = "User" // Default to "User" if LOGNAME is not set
	}
	fmt.Printf("Hello %s!\n", logName)
	fmt.Println()

	// Display current date and time
	fmt.Println("Today's date and time:")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")) // Format the date and time

	fmt.Println() // Print an empty line

	// Define variables
	myNum := 50
	myDay := "Sunday"

	// Display variable values
	fmt.Printf("The value of my_num is %d\n", myNum)
	fmt.Printf("The value of my_day is %s\n", myDay)
	fmt.Println()
	fmt.Println("SCRIPT FINISHED!!")
}

// Function to clear the terminal window
func clearTerminal() {
	cmd := exec.Command("clear") // Use "cls" on Windows
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing the terminal:", err)
	}
}

