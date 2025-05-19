package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open or create the output.txt file for writing
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer outputFile.Close()

	// Open the /proc/cpuinfo file for reading
	cpuInfoFile, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error opening /proc/cpuinfo:", err)
		return
	}
	defer cpuInfoFile.Close()

	// Copy the contents of /proc/cpuinfo to output.txt
	_, err = io.Copy(outputFile, cpuInfoFile)
	if err != nil {
		fmt.Println("Error copying content to the file:", err)
		return
	}

	fmt.Println("Contents of /proc/cpuinfo have been written to output.txt")
}

