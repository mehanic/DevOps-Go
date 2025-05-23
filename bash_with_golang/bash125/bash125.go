package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"

    "flag"
)

// Function to backup a file
func backupFile(filePath string) {
    if _, err := os.Stat(filePath); err == nil {
        backupPath := filepath.Join("/var/tmp", fmt.Sprintf("%s.%s.%d", filepath.Base(filePath), time.Now().Format("2006-01-02"), os.Getpid()))
        fmt.Printf("Backing up %s to %s\n", filePath, backupPath)
        err = os.Rename(filePath, backupPath) // Using rename as a simple way to move the file
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error backing up file: %v\n", err)
            os.Exit(1)
        }
    }
}

// Function to exit if a file does not exist
func exitIfFileDoesNotExist(filePath string) {
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        fmt.Printf("%s does not exist. Exiting.\n", filePath)
        os.Exit(1)
    }
}

// Function to create slave device file
func makeSlaveDeviceFile(sDevice, sFile, bondD string) {
    content := fmt.Sprintf(`DEVICE=%s
BOOTPROTO=none
ONBOOT=yes
MASTER=%s
SLAVE=yes
USERCTL=no
NM_CONTROLLED=no
`, sDevice, bondD)
    err := ioutil.WriteFile(sFile, []byte(content), 0644)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", sFile, err)
        os.Exit(1)
    }
}

func main() {
    // Check if the program is running as root
    if os.Geteuid() != 0 {
        fmt.Fprintln(os.Stderr, "Please run with sudo or as root.")
        os.Exit(1)
    }

    // Set defaults
    sourceSlaveDevice := "eth0"
    partnerSlaveDevice := "eth1"
    bondDevice := "bond0"
    restartNetworking := false

    // Parse command-line flags
    flag.StringVar(&sourceSlaveDevice, "s", sourceSlaveDevice, "Source NIC")
    flag.StringVar(&partnerSlaveDevice, "p", partnerSlaveDevice, "Partner NIC")
    flag.StringVar(&bondDevice, "b", bondDevice, "Bond device")
    flag.BoolVar(&restartNetworking, "r", restartNetworking, "Restart networking after the bond has been configured")
    flag.Parse()

    // Configuration file variables
    ifcfgPath := "/etc/sysconfig/network-scripts"
    bondDeviceFile := filepath.Join(ifcfgPath, fmt.Sprintf("ifcfg-%s", bondDevice))
    sourceSlaveDeviceFile := filepath.Join(ifcfgPath, fmt.Sprintf("ifcfg-%s", sourceSlaveDevice))
    partnerSlaveDeviceFile := filepath.Join(ifcfgPath, fmt.Sprintf("ifcfg-%s", partnerSlaveDevice))
    bondingOpts := "mode=1 miimon=50 updelay=100"

    // Check for RHEL 5 and add alias if needed
    if strings.Contains(string(readFile("/etc/redhat-release")), " 5.") {
        aliasLine := fmt.Sprintf("alias %s bonding", bondDevice)
        if !fileContains("/etc/modprobe.conf", aliasLine) {
            backupFile("/etc/modprobe.conf")
            appendToFile("/etc/modprobe.conf", aliasLine+"\n")
        }
    }

    // Create the basic config for the bonded device
    backupFile(bondDeviceFile)
    content := fmt.Sprintf(`DEVICE=%s
ONBOOT=yes
BONDING_OPTS="%s"
NM_CONTROLLED=no
`, bondDevice, bondingOpts)
    ioutil.WriteFile(bondDeviceFile, []byte(content), 0644)

    // Append source device configuration to bond device
    backupFile(sourceSlaveDeviceFile)
    exitIfFileDoesNotExist(sourceSlaveDeviceFile)
    appendDeviceConfig(sourceSlaveDeviceFile, bondDeviceFile)

    fmt.Printf("Created %s\n", bondDeviceFile)

    // Create configuration for slave devices
    makeSlaveDeviceFile(sourceSlaveDevice, sourceSlaveDeviceFile, bondDevice)
    backupFile(partnerSlaveDeviceFile)
    makeSlaveDeviceFile(partnerSlaveDevice, partnerSlaveDeviceFile, bondDevice)

    // Restart networking if requested
    if restartNetworking {
        cmd := exec.Command("service", "network", "restart")
        err := cmd.Run()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Failed to restart networking: %v\n", err)
        }
    }
}

// Read file content
func readFile(filePath string) []byte {
    content, err := ioutil.ReadFile(filePath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filePath, err)
        os.Exit(1)
    }
    return content
}

// Check if a file contains a specific line
func fileContains(filePath, searchString string) bool {
    content := readFile(filePath)
    return strings.Contains(string(content), searchString)
}

// Append a line to a file
func appendToFile(filePath, line string) {
    f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filePath, err)
        os.Exit(1)
    }
    defer f.Close()
    if _, err := f.WriteString(line); err != nil {
        fmt.Fprintf(os.Stderr, "Error writing to file %s: %v\n", filePath, err)
        os.Exit(1)
    }
}

// Append the necessary configuration from source device to bond device
func appendDeviceConfig(sourceFile, bondFile string) {
    content, err := ioutil.ReadFile(sourceFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", sourceFile, err)
        os.Exit(1)
    }
    // Filter out unwanted lines
    lines := strings.Split(string(content), "\n")
    for _, line := range lines {
        if !strings.Contains(line, "DEVICE") && 
           !strings.Contains(line, "NM_CONTROLLED") && 
           !strings.Contains(line, "HWADDR") && 
           !strings.Contains(line, "UUID") && 
           !strings.Contains(line, "ONBOOT") && 
           !strings.Contains(line, "TYPE") {
            appendToFile(bondFile, line+"\n")
        }
    }
}

