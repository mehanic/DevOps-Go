package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// checkProcess checks if a given process name exists
func checkProcess(processName string) (bool, error) {
	if processName == "" {
		return false, nil
	}

	// Execute the ps command to list processes
	cmd := exec.Command("ps", "-ef")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	// Convert the output to a string and split it into lines
	processList := string(output)
	lines := strings.Split(processList, "\n")

	// Count the number of occurrences of the process name
	count := 0
	for _, line := range lines {
		if strings.Contains(line, processName) && !strings.Contains(line, "grep") {
			count++
		}
	}

	return count > 0, nil
}

func main() {
	processName := "mysql" // Change this to the process you want to check

	exists, err := checkProcess(processName)
	if err != nil {
		fmt.Printf("Error checking process: %v\n", err)
		return
	}

	if exists {
		// Code block for process exists
		fmt.Printf("Process '%s' exists.\n", processName)
	} else {
		// Code block for process not present
		fmt.Printf("Process '%s' does not exist.\n", processName)
	}
}

