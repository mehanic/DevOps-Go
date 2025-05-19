package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Get the LOGNAME environment variable
	logname := os.Getenv("LOGNAME")
	if logname == "" {
		logname = "User" // Fallback if LOGNAME is not set
	}

	// Print Hello message
	fmt.Printf("Hello %s\n", logname)

	// Get and print the current date
	dateCmd := exec.Command("date")
	dateOutput, err := dateCmd.Output()
	if err != nil {
		fmt.Println("Error getting date:", err)
		return
	}
	fmt.Printf("Today is %s", string(dateOutput))

	// Get and print the present working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting present working directory:", err)
		return
	}
	fmt.Printf("Your present working directory is %s\n", pwd)

	// Print Goodbye message
	fmt.Printf("Good-bye %s\n", logname)
}

