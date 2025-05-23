package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Step 1: Copy /etc/hosts to hosts2
	srcFile, err := os.Open("/etc/hosts")
	if err != nil {
		fmt.Println("Error opening /etc/hosts:", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create("hosts2")
	if err != nil {
		fmt.Println("Error creating hosts2:", err)
		return
	}
	defer destFile.Close()

	// Copy contents of /etc/hosts to hosts2
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		fmt.Println("Error copying /etc/hosts to hosts2:", err)
		return
	}

	// Step 2: Remove lines starting with '#' from hosts2 and write to hosts3
	srcFile, err = os.Open("hosts2")
	if err != nil {
		fmt.Println("Error opening hosts2:", err)
		return
	}
	defer srcFile.Close()

	hosts3File, err := os.Create("hosts3")
	if err != nil {
		fmt.Println("Error creating hosts3:", err)
		return
	}
	defer hosts3File.Close()

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(hosts3File)

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			writer.WriteString(line + "\n")
		}
	}

	// Ensure buffered data is written to the file
	if err := writer.Flush(); err != nil {
		fmt.Println("Error writing to hosts3:", err)
		return
	}

	// Step 3: Read first two lines of hosts3 and write specific fields to hostsfinal
	hosts3File, err = os.Open("hosts3")
	if err != nil {
		fmt.Println("Error opening hosts3:", err)
		return
	}
	defer hosts3File.Close()

	hostsFinalFile, err := os.Create("hostsfinal")
	if err != nil {
		fmt.Println("Error creating hostsfinal:", err)
		return
	}
	defer hostsFinalFile.Close()

	scanner = bufio.NewScanner(hosts3File)
	writer = bufio.NewWriter(hostsFinalFile)

	// Read and write the first two lines
	if scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			address1, name1 := fields[0], fields[1]
			writer.WriteString(fmt.Sprintf("%s %s\n", name1, address1))
		}
	}

	if scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			address2, name2 := fields[0], fields[1]
			writer.WriteString(fmt.Sprintf("%s %s\n", name2, address2))
		}
	}

	// Flush final content to hostsfinal
	if err := writer.Flush(); err != nil {
		fmt.Println("Error writing to hostsfinal:", err)
		return
	}

	fmt.Println("Processing complete.")
}

