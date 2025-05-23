package main

import (
	"fmt"
	"os/exec"
//	"strings"
	"time"
)

// Function to print the current system status.
func printSystemStatus() {
	// Get the current datetime.
	fmt.Println("Date:", time.Now().Format("2006-01-02 15:04:05"))

	// Get the CPU usage.
	cpuUsage, err := exec.Command("sh", "-c", "top -bn1 | grep 'Cpu' | awk '{print $2}'").Output()
	if err == nil {
		fmt.Printf("CPU in use: %s", string(cpuUsage))
	}

	// Get the memory usage.
	memUsage, err := exec.Command("sh", "-c", "free -h | grep Mem | awk '{print $3}'").Output()
	if err == nil {
		fmt.Printf("Memory in use: %s", string(memUsage))
	}

	// Get the disk space available on "/".
	diskSpace, err := exec.Command("sh", "-c", "df -k / | grep / | awk '{print $4}'").Output()
	if err == nil {
		fmt.Printf("Disk space available on /: %s", string(diskSpace))
	}

	// Add an extra newline for readability.
	fmt.Println()
}

func main() {
	// Print the system status five times, with a 5-second interval.
	for i := 0; i < 5; i++ {
		printSystemStatus()
		time.Sleep(5 * time.Second)
	}
}

