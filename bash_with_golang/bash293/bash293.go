package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Take first process snapshot and save it to p_snapshot1.txt
	err := takeProcessSnapshot("p_snapshot1.txt")
	if err != nil {
		fmt.Println("Error taking process snapshot:", err)
		return
	}

	// Count and print the number of running processes at snapshot1
	countProcesses("p_snapshot1.txt", "Running process count at snapshot1")

	// Create a new process: tail -f /dev/null & and get its PID
	fmt.Print("Create a new process with pid = ")
	cmd := exec.Command("tail", "-f", "/dev/null")
	err = cmd.Start() // Start process in the background
	if err != nil {
		fmt.Println("Error creating new process:", err)
		return
	}
	fmt.Println(cmd.Process.Pid) // Print the new process's PID

	// Take the second process snapshot
	err = takeProcessSnapshot("p_snapshot2.txt")
	if err != nil {
		fmt.Println("Error taking second process snapshot:", err)
		return
	}

	// Count and print the number of running processes at snapshot2
	countProcesses("p_snapshot2.txt", "Running process count at snapshot2")

	// Show the difference between the two snapshots
	fmt.Println("\nDiff between two snapshots:")
	diffFiles("p_snapshot1.txt", "p_snapshot2.txt")
}

// takeProcessSnapshot captures the process list and stores it in a file
func takeProcessSnapshot(filename string) error {
	cmd := exec.Command("ps", "-A", "-o", "pid", "-o", "command")
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	return os.WriteFile(filename, output, 0644)
}

// countProcesses prints the number of lines (processes) in the snapshot file
func countProcesses(filename, message string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	fmt.Printf("%s: %d\n", message, lineCount)
}

// diffFiles shows the difference between two files
func diffFiles(file1, file2 string) {
	cmd := exec.Command("diff", file1, file2)
	output, err := cmd.Output()
	if err != nil && !strings.Contains(err.Error(), "exit status 1") { // diff returns status 1 if there's a difference
		fmt.Println("Error running diff:", err)
		return
	}
	fmt.Println(string(output))
}

