package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "regexp"
    "strings"
)

func main() {
    // Check if the directory "/sys/class/fc_host" exists
    dir := "/sys/class/fc_host"
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        fmt.Println("Directory does not exist:", dir)
        return
    }

    // Read all port_name files from the directory
    files, err := filepath.Glob(filepath.Join(dir, "*", "port_name"))
    if err != nil {
        fmt.Println("Error reading port_name files:", err)
        return
    }

    var portNames []string

    // Process each file
    for _, file := range files {
        data, err := ioutil.ReadFile(file)
        if err != nil {
            fmt.Println("Error reading file:", file, err)
            continue
        }

        // Trim the data and remove the "0x" prefix
        trimmedData := strings.TrimSpace(string(data))
        if strings.HasPrefix(trimmedData, "0x") {
            trimmedData = strings.TrimPrefix(trimmedData, "0x")
        }

        // Format the port name
        formatted := formatPortName(trimmedData)
        portNames = append(portNames, formatted)
    }

    // Print the formatted port names
    for _, name := range portNames {
        fmt.Println(name)
    }
}

// formatPortName adds colons between every two characters of the input string
func formatPortName(portName string) string {
    var sb strings.Builder
    re := regexp.MustCompile(`(.{2})`)
    matches := re.FindAllString(portName, -1)
    for _, match := range matches {
        sb.WriteString(match + ":")
    }
    return strings.TrimSuffix(sb.String(), ":") // Remove trailing colon
}

