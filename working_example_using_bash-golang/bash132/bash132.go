package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

// getScheduler retrieves unique I/O schedulers from the /sys/block directory
func getScheduler() ([]string, error) {
    schedulers := make(map[string]struct{})
    blockDir := "/sys/block"

    // Check if the directory exists
    if _, err := os.Stat(blockDir); os.IsNotExist(err) {
        return nil, fmt.Errorf("directory %s does not exist", blockDir)
    }

    // Walk through the /sys/block directory
    err := filepath.Walk(blockDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Look for queue/scheduler files
        if info.IsDir() || !strings.HasSuffix(path, "scheduler") {
            return nil
        }

        // Open and read the scheduler file
        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        if scanner.Scan() {
            // Extract the scheduler name, avoiding "none"
            line := scanner.Text()
            if line != "none" {
                parts := strings.Split(line, "[")
                if len(parts) > 1 {
                    scheduler := strings.Split(parts[1], "]")[0]
                    schedulers[scheduler] = struct{}{} // Use a set to avoid duplicates
                }
            }
        }
        return nil
    })

    if err != nil {
        return nil, err
    }

    // Convert map keys to a slice
    var uniqueSchedulers []string
    for scheduler := range schedulers {
        uniqueSchedulers = append(uniqueSchedulers, scheduler)
    }

    // Sort the slice
    sort.Strings(uniqueSchedulers)
    return uniqueSchedulers, nil
}

// main function
func main() {
    schedulers, err := getScheduler()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the unique schedulers
    for _, scheduler := range schedulers {
        fmt.Println(scheduler)
    }
}

