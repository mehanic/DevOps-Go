package main

import (
	"bytes"
	"fmt"
//	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Function to check available space and retrieve size of tarball contents
func checkDiskSpaceAndTarballSize(tarball string) (available, used, total int64, compressedSize float64, err error) {
	// Get disk space information
	cmd := exec.Command("df", "--output=size,avail,used", "/home", "-B", "1M")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// Read output and parse values
	lines := strings.Split(out.String(), "\n")
	if len(lines) < 2 {
		return 0, 0, 0, 0, fmt.Errorf("failed to get disk space information")
	}

	// Parse available and used space
	fields := strings.Fields(lines[1])
	total, _ = strconv.ParseInt(fields[0], 10, 64)
	available, _ = strconv.ParseInt(fields[1], 10, 64)
	used, _ = strconv.ParseInt(fields[2], 10, 64)

	// Get the size of the contents of the tarball
	cmd = exec.Command("tar", "tzvf", tarball)
	out.Reset()
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	// Calculate compressed size
	tarContents := strings.Split(out.String(), "\n")
	var totalSize float64
	for _, line := range tarContents {
		if len(line) > 0 {
			parts := strings.Fields(line)
			if len(parts) >= 3 {
				size, _ := strconv.ParseFloat(parts[2], 64)
				totalSize += size
			}
		}
	}

	compressedSize = totalSize
	return available, used, total, compressedSize, nil
}

func main() {
	tarball := "archive.tar.gz"

	// Check disk space and tarball size
	available, used, total, compressedSize, err := checkDiskSpaceAndTarballSize(tarball)
	if err != nil {
		log.Fatalf("Error retrieving disk space or tarball size: %v", err)
	}

	fmt.Println("First inspection - is there enough space?")
	if available < int64(compressedSize) {
		fmt.Println("There is not enough space to decompress the binary")
		os.Exit(1)
	} else {
		fmt.Println("Seems we have enough space on first inspection - continuing")
		spaceLeft := total - used
		fmt.Printf("Space left: %d MB\n", spaceLeft)
	}

	fmt.Println("Safety check - do we have at least double the space?")
	doubleAvailable := available * 2
	fmt.Printf("Double - good question? %d MB\n", doubleAvailable)

	// Check if available space is greater than zero after dividing by 2
	if float64(available)/2 > 0 {
		fmt.Println("Sppppppaaaaccee (imagine zombies)")
	}

	// Untar the file
	cmd := exec.Command("tar", "xzvf", tarball)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error decompressing the tarball!")
		os.Exit(1)
	}

	fmt.Println("Decompressing tarball complete!")
	os.Exit(0)
}

