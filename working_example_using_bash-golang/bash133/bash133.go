package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "os/user"
    "regexp"
    "sort"
    "strings"
    "time"
)

// Function to check if the user is root
func isRoot() bool {
    currentUser, err := user.Current()
    if err != nil {
        return false
    }
    return currentUser.Uid == "0"
}

// Function to get the list of network interfaces
func getNics() ([]string, error) {
    cmd := exec.Command("ip", "link")
    output, err := cmd.Output()
    if err != nil {
        return nil, err
    }

    var nics []string
    scanner := bufio.NewScanner(strings.NewReader(string(output)))

    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, " ") || line == "" {
            continue
        }
        
        parts := strings.Split(line, ":")
        if len(parts) < 2 {
            continue
        }

        nic := strings.TrimSpace(parts[1])
        
        // Exclude unwanted NICs
        if matched, _ := regexp.MatchString(`^(bond|lo|usb|@|vnet|vir|br)`, nic); !matched {
            nics = append(nics, nic)
        }
    }

    sort.Strings(nics)
    return nics, nil
}

// Function to run tcpdump on the specified NIC with a timeout
func runTcpdump(nic string, timeoutSeconds int) {
    cmd := exec.Command("tcpdump", "-v", "-i", nic, "-s", "1500", "-c", "1", "ether[20:2] == 0x2000")

    // Create a pipe for reading command output
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Printf("Failed to create pipe: %v\n", err)
        return
    }

    // Start the command
    if err := cmd.Start(); err != nil {
        fmt.Printf("Failed to start tcpdump: %v\n", err)
        return
    }

    // Create a channel to handle timeout
    done := make(chan struct{})
    go func() {
        scanner := bufio.NewScanner(stdout)
        for scanner.Scan() {
            line := scanner.Text()
            if matched, _ := regexp.MatchString(`Device-ID|Port-ID|VLAN`, line); matched {
                // Print the last field and clean it
                fields := strings.Fields(line)
                if len(fields) > 0 {
                    fmt.Println(strings.Trim(fields[len(fields)-1], "'"))
                }
            }
        }
        done <- struct{}{}
    }()

    // Wait for command completion or timeout
    select {
    case <-time.After(time.Duration(timeoutSeconds) * time.Second):
        cmd.Process.Kill()
        fmt.Println("tcpdump timed out.")
    case <-done:
        // Command finished before timeout
    }

    // Wait for command to finish
    if err := cmd.Wait(); err != nil {
        fmt.Printf("tcpdump exited with error: %v\n", err)
    }
}

func main() {
    const timeoutSeconds = 60

    // Check if running as root
    if !isRoot() {
        fmt.Println("Please run with sudo or as root.")
        return
    }

    // Check for NIC arguments
    var nics []string
    if len(os.Args) > 1 {
        nics = os.Args[1:]
    } else {
        var err error
        nics, err = getNics()
        if err != nil {
            fmt.Printf("Failed to get NICs: %v\n", err)
            return
        }
    }

    // Loop through the NICs
    for _, nic := range nics {
        fmt.Println(nic)
        runTcpdump(nic, timeoutSeconds)
        fmt.Println()
    }
}

