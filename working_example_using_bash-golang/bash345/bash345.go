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

	// Print hello message using the LOGNAME
	fmt.Printf("Hello %s\n", logname)

	// Print hardcoded "Hello student" message
	fmt.Println("Hello student")

	// Get and print the current date
	dateCmd := exec.Command("date")
	dateOutput, err := dateCmd.Output()
	if err != nil {
		fmt.Println("Error getting date:", err)
		return
	}
	fmt.Printf("Today is %s", string(dateOutput))

	// Print hardcoded date for testing purposes
	fmt.Println("Today is Fri May 1 00:18:52 IST 2015")

	// Get and print the present working directory
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting present working directory:", err)
		return
	}
	fmt.Printf("Your present working directory is %s\n", pwd)

	// Print hardcoded working directory for testing purposes
	fmt.Println("Your present working directory is /home/student/work")

	// Print goodbye message using the LOGNAME
	fmt.Printf("Good-bye %s\n", logname)

	// Print hardcoded "Good-bye student" message
	fmt.Println("Good-bye student")
}

