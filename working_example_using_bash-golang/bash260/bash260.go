package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
//	"path/filepath"
)

func main() {
	// Open file for reading (similar to exec 3< test.txt in bash)
	inputFile, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening test.txt:", err)
		return
	}
	defer inputFile.Close()

	// Open file for writing (similar to exec 4> output.txt in bash)
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output.txt:", err)
		return
	}
	defer outputFile.Close()

	// Read a single line from the input file (like `read -u 3 line`)
	scanner := bufio.NewScanner(inputFile)
	if scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line read from test.txt:", line)

		// Optionally, write the line to output.txt (or anything else)
		_, err = outputFile.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to output.txt:", err)
			return
		}
	}

	// Print current process ID (like `$$` in bash)
	myPID := os.Getpid()
	fmt.Printf("Process ID of the current process is %d\n", myPID)

	// List currently open file descriptors for the process (similar to `ls -l /proc/$my_pid/fd`)
	fmt.Println("Currently open files by this process:")
	cmd := exec.Command("ls", "-l", fmt.Sprintf("/proc/%d/fd", myPID))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error listing file descriptors:", err)
		return
	}

	// Files are automatically closed when we defer the Close calls above
}

