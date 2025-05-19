package main

import (
	"fmt"
//	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
//	"time"
)

func main() {
	// Find the latest db2diag log file
	db2DiagLog := findLatestDb2DiagLog("/db2")

	if db2DiagLog == "" {
		fmt.Println("No db2 diag log file found.")
		os.Exit(1)
	}

	// Show the log file being tailed
	fmt.Println(db2DiagLog)
	fmt.Println()

	// Tail the log file
	if err := tailFile(db2DiagLog); err != nil {
		log.Fatalf("Error tailing log file: %v", err)
	}
}

// findLatestDb2DiagLog finds the latest db2diag log file
func findLatestDb2DiagLog(baseDir string) string {
	var files []string

	// Walk through the directory structure
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.Contains(info.Name(), "db2diag") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the directory: %v", err)
	}

	// Sort files by modification time
	if len(files) == 0 {
		return ""
	}
	sort.Slice(files, func(i, j int) bool {
		fi, _ := os.Stat(files[i])
	 fj, _ := os.Stat(files[j])
		return fi.ModTime().Before(fj.ModTime())
	})

	// Return the last file in the sorted list (the latest)
	return files[len(files)-1]
}

// tailFile simulates the tail -f command for the specified log file
func tailFile(filename string) error {
	// Run the tail command
	cmd := exec.Command("tail", "-f", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

