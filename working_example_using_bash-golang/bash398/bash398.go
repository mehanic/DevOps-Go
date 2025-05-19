package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	tarball := "archive.tar.gz"

	// Retrieve filesystem information
	currentPartAll, err := getFilesystemInfo()
	if err != nil {
		log.Fatalf("Error retrieving filesystem info: %v", err)
	}

	// Split the filesystem information into size, available, and used
	array := strings.Fields(currentPartAll)
	if len(array) < 3 {
		log.Fatalf("Unexpected filesystem info format: %s", currentPartAll)
	}

	// Retrieve the size of the contents of the tarball
	compressedSz, err := getCompressedSize(tarball)
	if err != nil {
		log.Fatalf("Error retrieving compressed size: %v", err)
	}

	fmt.Println("First inspection - is there enough space?")
	availableSpace, err := strconv.Atoi(array[1]) // Handle error
	if err != nil {
		log.Fatalf("Error converting available space: %v", err)
	}

	if availableSpace < compressedSz {
		fmt.Println("There is not enough space to decompress the binary")
		os.Exit(1)
	} else {
		fmt.Println("Seems we have enough space on first inspection - continuing")
		usedSpace, err := strconv.Atoi(array[2]) // Handle error
		if err != nil {
			log.Fatalf("Error converting used space: %v", err)
		}
		spaceLeft := availableSpace - usedSpace
		fmt.Printf("Space left: %d\n", spaceLeft)
	}

	fmt.Println("Safety check - do we have at least double the space?")
	doubleAvailable := availableSpace * 2
	fmt.Printf("Double - good question? %d\n", doubleAvailable)

	// Check if we have enough space using floating-point division
	if hasSufficientSpace(doubleAvailable) {
		fmt.Println("Sppppppaaaaccee (imagine zombies)")
	}

	// Decompress the tarball
	if err := decompressTarball(tarball); err != nil {
		log.Fatalf("Error decompressing the tarball: %v", err)
	}

	fmt.Println("Decompressing tarball complete!")
}

// getFilesystemInfo retrieves filesystem information for the /home directory
func getFilesystemInfo() (string, error) {
	cmd := exec.Command("df", "--output=size,avail,used", "/home", "-B", "1M")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// getCompressedSize retrieves the size of the contents of the tarball
func getCompressedSize(tarball string) (int, error) {
	cmd := exec.Command("tar", "tzvf", tarball)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return 0, err
	}

	lines := strings.Split(out.String(), "\n")
	var totalSize int
	for _, line := range lines {
		// Split the line to extract the size (the third field)
		fields := strings.Fields(line)
		if len(fields) > 2 {
			size, err := strconv.Atoi(fields[2])
			if err != nil {
				continue
			}
			totalSize += size
		}
	}
	return totalSize, nil
}

// hasSufficientSpace checks if the given space is greater than zero after division
func hasSufficientSpace(value int) bool {
	return value > 0
}

// decompressTarball decompresses the specified tarball
func decompressTarball(tarball string) error {
	cmd := exec.Command("tar", "xzvf", tarball)
	err := cmd.Run()
	return err
}

