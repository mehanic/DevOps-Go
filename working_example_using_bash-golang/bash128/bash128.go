package main

import (
//    "flag"
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strings"
)

// Function to ping a host and return whether it is up or down
func pingHost(host string) (bool, error) {
    var cmd *exec.Cmd

    // Determine the correct ping command based on the OS
    if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
        cmd = exec.Command("ping", "-c", "1", "-W", "1", host) // Linux / Mac
    } else if runtime.GOOS == "solaris" {
        cmd = exec.Command("ping", "-t", "10", host) // Solaris
    } else {
        return false, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
    }

    // Run the ping command and get the output
    output, err := cmd.CombinedOutput()
    if err != nil {
        return false, err
    }

    // Check if the output indicates success
    if strings.Contains(string(output), "1 packets transmitted, 1 received") || 
       strings.Contains(string(output), "is alive") { // Solaris might have different output
        return true, nil
    }

    return false, nil
}

func usage() {
    fmt.Fprintln(os.Stderr, "Usage: ping-host HOST [HOSTN]...")
    fmt.Fprintln(os.Stderr, "Returns 0 if the hosts respond to ping or 1 if any of them fail to respond.")
    os.Exit(1)
}

func main() {
    // Check for at least one argument
    if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
        usage()
    }

    // Assume all hosts are reachable initially
    returnCode := 0

    // Loop through the hosts provided on the command line
    for _, host := range os.Args[1:] {
        reachable, err := pingHost(host)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error pinging %s: %v\n", host, err)
            returnCode = 1
            continue
        }
        if !reachable {
            fmt.Printf("%s down\n", host)
            returnCode = 1
        }
    }

    os.Exit(returnCode)
}

