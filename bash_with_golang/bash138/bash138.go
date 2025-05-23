package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
//	"strconv"
	"strings"
	"time"
)

func main() {
	// Check if the program is running as root
	if os.Geteuid() != 0 {
		log.Fatal("Please run with sudo or as root.")
	}

	testFile := "touch-test-file"

	start := time.Now()
	startSeconds := time.Now().Unix()

	// Get list of mount points excluding tmpfs
	mounts, err := getMountPoints()
	if err != nil {
		log.Fatalf("Error getting mount points: %v", err)
	}

	for _, mount := range mounts {
		testFileOnMount := filepath.Join(mount, testFile)
		fmt.Printf("%s - Touching %s\n", time.Now().Format(time.RFC3339), testFileOnMount)

		// Create the test file
		if err := touchFile(testFileOnMount); err != nil {
			log.Printf("Failed to touch %s: %v", testFileOnMount, err)
			continue
		}

		// Remove the test file
		if err := os.Remove(testFileOnMount); err != nil {
			log.Printf("Failed to remove %s: %v", testFileOnMount, err)
			continue
		}

		fmt.Printf("%s - Removed %s\n", time.Now().Format(time.RFC3339), testFileOnMount)
	}

	end := time.Now()
	endSeconds := time.Now().Unix()
	totalSeconds := endSeconds - startSeconds

	fmt.Println()
	fmt.Printf("Start: %s\n", start.Format(time.RFC3339))
	fmt.Printf("End:   %s\n", end.Format(time.RFC3339))
	fmt.Printf("Total: %d seconds\n", totalSeconds)
}

// getMountPoints returns a list of mount points excluding tmpfs
func getMountPoints() ([]string, error) {
	cmd := exec.Command("df", "-lP")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var mounts []string
	lines := strings.Split(string(output), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "Filesystem") || strings.Contains(line, "tmpfs") {
			continue // Skip headers and tmpfs
		}

		// Split line into fields
		fields := strings.Fields(line)
		if len(fields) > 0 {
			mounts = append(mounts, fields[len(fields)-1]) // The last field is the mount point
		}
	}

	return mounts, nil
}

// touchFile creates a file at the specified path
func touchFile(path string) error {
	_, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	return err
}

