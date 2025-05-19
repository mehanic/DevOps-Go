package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
//    "regexp"
    "strconv"
    "strings"
)

// executeCommand runs a shell command and returns its output as a string
func executeCommand(name string, args ...string) (string, error) {
    cmd := exec.Command(name, args...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}

// getMemoryInfo calculates total memory in GB
func getMemoryInfo() (int, error) {
    output, err := executeCommand("dmidecode", "--type", "memory")
    if err != nil {
        return 0, err
    }

    var totalMemoryMB int
    scanner := bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "Size:") {
            fields := strings.Fields(line)
            if len(fields) > 1 {
                size, err := strconv.Atoi(fields[1])
                if err == nil {
                    totalMemoryMB += size
                }
            }
        }
    }
    return totalMemoryMB / 1024, nil // Convert to GB
}

// getCPUInfo gathers CPU information
func getCPUInfo() (map[string]int, map[string]int, map[string]int, error) {
    output, err := executeCommand("cat", "/proc/cpuinfo")
    if err != nil {
        return nil, nil, nil, err
    }

    cpuModel := make(map[string]int)
    totalCores := make(map[string]int)
    totalThreads := make(map[string]int)

    scanner := bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "model name") {
            model := strings.TrimSpace(strings.Split(line, ":")[1])
            cpuModel[model]++
        } else if strings.HasPrefix(line, "physical id") {
            physicalID := strings.TrimSpace(strings.Split(line, ":")[1])
            cores, err := totalCores[physicalID]
            if err {
                cores = 0
            }
            cores++
            totalCores[physicalID] = cores
            totalThreads[physicalID]++
        }
    }

    return cpuModel, totalCores, totalThreads, nil
}

// getStorageInfo calculates storage assigned to volume groups in GB
func getStorageInfo() (int, error) {
    output, err := executeCommand("vgs", "--units", "g", "--noheadings")
    if err != nil {
        return 0, err
    }

    var totalStorageGB int
    scanner := bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        if len(fields) > 5 {
            if size, err := strconv.Atoi(strings.Split(fields[5], ".")[0]); err == nil {
                totalStorageGB += size
            }
        }
    }
    return totalStorageGB, nil
}

// getNetworkInfo retrieves network interfaces and IP addresses
func getNetworkInfo() (string, string, error) {
    // Get network interfaces
    output, err := executeCommand("netstat", "-i")
    if err != nil {
        return "", "", err
    }
    
    interfaces := make([]string, 0)
    scanner := bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        line := scanner.Text()
        if !strings.Contains(line, "Iface") && !strings.Contains(line, "Interface") {
            fields := strings.Fields(line)
            if len(fields) > 0 && fields[0] != "lo" {
                interfaces = append(interfaces, fields[0])
            }
        }
    }

    // Get IP addresses
    output, err = executeCommand("ip", "-4", "-o", "addr")
    if err != nil {
        return "", "", err
    }

    ips := make([]string, 0)
    scanner = bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        if len(fields) > 3 {
            ip := strings.Split(fields[3], "/")[0]
            if ip != "127.0.0.1" {
                ips = append(ips, ip)
            }
        }
    }

    return strings.Join(interfaces, ", "), strings.Join(ips, ", "), nil
}

// main function
func main() {
    // Check if running as root
    if os.Geteuid() != 0 {
        fmt.Println("Please run with sudo or as root.")
        os.Exit(1)
    }

    fieldSeparator := ": "
    if len(os.Args) > 1 && os.Args[1] == "-c" {
        fieldSeparator = ", "
    }

    // Host name
    hostName, _ := executeCommand("uname", "-n")
    fmt.Printf("Hostname%s%s", fieldSeparator, strings.TrimSpace(hostName))
    
    fqdn, _ := executeCommand("hostname", "-f")
    fmt.Printf("\nFQDN%s%s", fieldSeparator, strings.TrimSpace(fqdn))

    // System model, product, serial
    dmidecodeOutput, _ := executeCommand("dmidecode", "-t", "system")
    scanner := bufio.NewScanner(strings.NewReader(dmidecodeOutput))
    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "Manufacturer:") || strings.Contains(line, "Product Name:") || strings.Contains(line, "Serial Number:") {
            cleanedLine := strings.TrimSpace(line)
            cleanedLine = strings.Replace(cleanedLine, ": ", fieldSeparator, 1)
            fmt.Println(cleanedLine)
        }
    }

    // OS, kernel, platform
    redhatRelease, _ := executeCommand("cat", "/etc/redhat-release")
    fmt.Printf("Release%s%s", fieldSeparator, strings.TrimSpace(redhatRelease))
    
    kernelRelease, _ := executeCommand("uname", "-r")
    fmt.Printf("\nKernel Release%s%s", fieldSeparator, strings.TrimSpace(kernelRelease))

    architecture, _ := executeCommand("uname", "-m")
    fmt.Printf("\nArchitecture%s%s", fieldSeparator, strings.TrimSpace(architecture))

    // Memory in GB
    memoryInGB, err := getMemoryInfo()
    if err == nil {
        fmt.Printf("\nMemory in GB%s%d", fieldSeparator, memoryInGB)
    }

    // CPU info
    cpuModel, totalCores, totalThreads, err := getCPUInfo()
    if err == nil {
        for model := range cpuModel {
            fmt.Printf("\nCPU Model%s%s", fieldSeparator, model)
        }
        for physicalID, cores := range totalCores {
            threads := totalThreads[physicalID]
            fmt.Printf("\nCPU %s Cores%s%d", physicalID, fieldSeparator, cores)
            fmt.Printf("\nCPU %s Threads%s%d", physicalID, fieldSeparator, threads)
        }
        fmt.Printf("\nTotal Physical CPUs%s%d", fieldSeparator, len(totalCores))
        totalCoreCount := 0
        for _, cores := range totalCores {
            totalCoreCount += cores
        }
        fmt.Printf("\nTotal Cores%s%d", fieldSeparator, totalCoreCount)
        fmt.Printf("\nTotal Threads%s%d", fieldSeparator, totalThreads)
    }

    // Storage assigned to volume groups in GB
    storageGB, err := getStorageInfo()
    if err == nil {
        fmt.Printf("\nStorage assigned to volume groups in GB%s%d", fieldSeparator, storageGB)
    }

    // Network interfaces and IP Addresses
    interfaces, ips, err := getNetworkInfo()
    if err == nil {
        fmt.Printf("\nNetwork interfaces%s%s", fieldSeparator, interfaces)
        fmt.Printf("\nIP Addresses%s%s", fieldSeparator, ips)
    }

    // Timezone
    timezone, _ := executeCommand("date", "+%Z")
    fmt.Printf("\nTimezone%s%s\n", fieldSeparator, strings.TrimSpace(timezone))
}

