package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// Clear the console
	clearConsole()

	// Print script beginning message
	fmt.Println("SCRIPT BEGINS")

	// Get the LOGNAME environment variable
	logname := os.Getenv("LOGNAME")
	if logname == "" {
		logname = "User" // Fallback if LOGNAME is not set
	}
	fmt.Printf("Hello %s!\n\n", logname)

	// Display current date and time
	fmt.Println("Today's date and time:")
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println()

	// Define variables
	myNum := 50
	myDay := "Sunday"

	// Print variable values
	fmt.Printf("The value of my_num is %d\n", myNum)
	fmt.Printf("The value of my_day is %s\n\n", myDay)

	// Print script finished message
	fmt.Println("SCRIPT FINISHED!!")
}

// clearConsole clears the terminal console screen
func clearConsole() {
	cmd := exec.Command("clear") // Use "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing console:", err)
	}
}

