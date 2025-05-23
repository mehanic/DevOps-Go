package main

import (
    "fmt"
    "os/exec"
    "regexp"
    "strconv"
    "strings"
)

// round function rounds the given floating-point number to the nearest integer
func round(value float64) int {
    integer := int(value)
    decimal := int((value - float64(integer)) * 10) // get the first decimal place
    if decimal > 4 {
        integer++
    }
    return integer
}

// getTotalDiskSpaceUsed retrieves the total disk space used in KB
func getTotalDiskSpaceUsed() (int, error) {
    cmd := exec.Command("df", "-lkP")
    output, err := cmd.Output()
    if err != nil {
        return 0, err
    }

    totalKB := 0
    lines := strings.Split(string(output), "\n")

    // Regular expression to match the output lines
    re := regexp.MustCompile(`^\S+\s+\S+\s+(\d+)\s+\S+\s+\S+\s+(\S+)$`)

    for _, line := range lines {
        if line == "" || strings.Contains(line, "Used") {
            continue // skip empty lines and header
        }
        match := re.FindStringSubmatch(line)
        if len(match) > 1 {
            if value, err := strconv.Atoi(match[1]); err == nil {
                totalKB += value
            }
        }
    }
    return totalKB, nil
}

// formatDiskSpace formats the disk space in a human-readable format
func formatDiskSpace(kb int) string {
    var formatted string
    mb := round(float64(kb) / 1024)
    gb := round(float64(kb) / (1024 * 1024))
    tb := round(float64(kb) / (1024 * 1024 * 1024))

    if tb > 0 {
        formatted = fmt.Sprintf("%dT", tb)
    } else if gb > 0 {
        formatted = fmt.Sprintf("%dG", gb)
    } else if mb > 0 {
        formatted = fmt.Sprintf("%dM", mb)
    } else {
        formatted = fmt.Sprintf("%dK", kb)
    }
    return formatted
}

func main() {
    // Get total disk space used
    totalKB, err := getTotalDiskSpaceUsed()
    if err != nil {
        fmt.Printf("Error retrieving disk space: %v\n", err)
        return
    }

    // Format and display the output
    hostname, _ := exec.Command("uname", "-n").Output()
    totalDiskSpaceUsed := formatDiskSpace(totalKB)
    fmt.Printf("%s\t%s\n", strings.TrimSpace(string(hostname)), totalDiskSpaceUsed)
}

