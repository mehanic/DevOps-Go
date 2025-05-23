package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strconv"
    "strings"
)

// executeCommand runs a shell command and returns its output as a string
func executeCommand(name string, args ...string) (string, error) {
    cmd := exec.Command(name, args...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}

// round rounds a floating-point number based on its decimal value
func round(value float64) int {
    integerPart := int(value)
    decimalPart := value - float64(integerPart)

    if decimalPart > 0.4 {
        return integerPart + 1
    }
    return integerPart
}

// getVolumeGroups retrieves the list of volume groups to process
func getVolumeGroups(args []string) ([]string, error) {
    if len(args) > 0 {
        return args, nil
    }

    output, err := executeCommand("vgs", "--noheadings")
    if err != nil {
        return nil, err
    }

    var volumeGroups []string
    scanner := bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        if len(fields) > 0 {
            volumeGroups = append(volumeGroups, fields[0])
        }
    }
    return volumeGroups, scanner.Err()
}

// getDiskUsage retrieves disk usage for a specified volume group
func getDiskUsage(vg string) (int64, error) {
    output, err := executeCommand("df", "-kP")
    if err != nil {
        return 0, err
    }

    totalKB := int64(0)
    scanner := bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, vg+"-") {
            fields := strings.Fields(line)
            if len(fields) > 2 {
                usedKB, err := strconv.ParseInt(fields[2], 10, 64)
                if err == nil {
                    totalKB += usedKB
                }
            }
        }
    }
    return totalKB, scanner.Err()
}

// formatSize converts kilobytes to the largest human-readable format
func formatSize(kb int64) (string, string) {
    const (
        KB_PER_MB = 1024
        MB_PER_GB = 1024
        GB_PER_TB = 1024
    )

    if tb := kb / (KB_PER_MB * MB_PER_GB * GB_PER_TB); tb > 0 {
        return strconv.FormatInt(tb, 10), "T"
    } else if gb := kb / (KB_PER_MB * MB_PER_GB); gb > 0 {
        return strconv.FormatInt(gb, 10), "G"
    } else if mb := kb / KB_PER_MB; mb > 0 {
        return strconv.FormatInt(mb, 10), "M"
    }
    return strconv.FormatInt(kb, 10), "K"
}

// main function
func main() {
    // Check if running as root
    if os.Geteuid() != 0 {
        fmt.Println("Please run with sudo or as root.")
        os.Exit(1)
    }

    // Get the list of volume groups
    vgs, err := getVolumeGroups(os.Args[1:])
    if err != nil {
        fmt.Println("Error retrieving volume groups:", err)
        os.Exit(1)
    }

    // Create a temporary file to store results
    tmpFile, err := os.CreateTemp("", "vg_usage_*.txt")
    if err != nil {
        fmt.Println("Error creating temporary file:", err)
        os.Exit(1)
    }
    defer os.Remove(tmpFile.Name()) // Clean up the temp file afterwards

    // Loop through each volume group and calculate disk usage
    for _, vg := range vgs {
        totalKB, err := getDiskUsage(vg)
        if err != nil {
            fmt.Printf("Error retrieving disk usage for %s: %s\n", vg, err)
            continue
        }

        // Convert and format the size
        size, unit := formatSize(totalKB)
        _, err = fmt.Fprintf(tmpFile, "%s %s%s\n", vg, size, unit)
        if err != nil {
            fmt.Println("Error writing to temporary file:", err)
            os.Exit(1)
        }
    }

    // Read and print the results from the temporary file
    tmpFile.Seek(0, 0) // Reset file pointer to the beginning
    scanner := bufio.NewScanner(tmpFile)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading temporary file:", err)
    }
}

