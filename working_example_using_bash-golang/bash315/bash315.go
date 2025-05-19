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

	// Print greetings
	fmt.Printf("Hello %s, Have a nice day!\n", logname)

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Printf("You are working in directory %s.\n", cwd)

	// Get the machine name (hostname)
	cmd := exec.Command("uname", "-n")
	hostname, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}
	fmt.Printf("You are working on a machine called %s.", string(hostname))

	// List files in the current directory
	fmt.Println("\nList of files in your directory is:")
	cmd = exec.Command("ls", "-la")
	cmd.Stdout = os.Stdout // Redirect output to standard output
	cmd.Stderr = os.Stderr // Redirect error to standard error
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing ls command:", err)
	}

	// Print the current working directory
	fmt.Println(cwd)

	// Print system information
	cmd = exec.Command("uname", "-a")
	cmd.Stdout = os.Stdout // Redirect output to standard output
	cmd.Stderr = os.Stderr // Redirect error to standard error
	if err := cmd.Run(); err != nil {
		fmt.Println("Error executing uname command:", err)
	}

	// Print current time
	cmd = exec.Command("date", "+%T")
	currentTime, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting current time:", err)
		return
	}
	fmt.Printf("Bye for now %s. The time is %s!\n", logname, string(currentTime))
}

