package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Step 1: Run the command `grep -R "model name" /proc/cpuinfo | sort -u`
	grepCmd := exec.Command("grep", "-R", "model name", "/proc/cpuinfo")
	sortCmd := exec.Command("sort", "-u")

	grepOut, err := grepCmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating grep pipe:", err)
		return
	}

	sortCmd.Stdin = grepOut
	sortOut, err := sortCmd.Output()
	if err != nil {
		fmt.Println("Error running sort command:", err)
		return
	}

	// Close grepCmd stdout to allow sortCmd to finish
	if err := grepCmd.Start(); err != nil {
		fmt.Println("Error starting grep command:", err)
		return
	}
	if err := grepCmd.Wait(); err != nil {
		fmt.Println("Error waiting for grep command:", err)
		return
	}

	// Write the output of the command to /tmp/tmp1.txt
	tmpFile1, err := os.Create("/tmp/tmp1.txt")
	if err != nil {
		fmt.Println("Error creating tmp1.txt:", err)
		return
	}
	defer tmpFile1.Close()

	tmpFile1.Write(sortOut)

	// Step 2: Remove spaces using Go’s string manipulation (equivalent to `tr -d ' '`)
	tmpFile2, err := os.Create("/tmp/tmp2.txt")
	if err != nil {
		fmt.Println("Error creating tmp2.txt:", err)
		return
	}
	defer tmpFile2.Close()

	scanner := bufio.NewScanner(strings.NewReader(string(sortOut)))
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		tmpFile2.WriteString(line + "\n")
	}

	// Step 3: Extract CPU speed (equivalent to `cut -d '@' -f2`)
	tmpFile2.Seek(0, 0) // Rewind the file for reading
	scanner = bufio.NewScanner(tmpFile2)
	for scanner.Scan() {
		line := scanner.Text()
		if parts := strings.Split(line, "@"); len(parts) > 1 {
			fmt.Println("Processor speed:", parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from tmp2.txt:", err)
	}
}

