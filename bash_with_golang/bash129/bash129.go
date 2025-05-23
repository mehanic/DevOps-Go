package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "regexp"
    "strings"
)

// Function to execute a command and return its output
func executeCommand(name string, args ...string) (string, error) {
    cmd := exec.Command(name, args...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}

// Function to get the volume groups
func getVolumeGroups() ([]string, error) {
    output, err := executeCommand("vgs", "--noheadings")
    if err != nil {
        return nil, err
    }

    var volumeGroups []string
    scanner := bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        volumeGroups = append(volumeGroups, strings.Fields(scanner.Text())[0])
    }
    return volumeGroups, nil
}

// Function to generate multipath aliases
func generateMultipathAliases(multipathConfFile string) error {
    if _, err := os.Stat(multipathConfFile); os.IsNotExist(err) {
        return fmt.Errorf("%s not found. Exiting.", multipathConfFile)
    }

    // Run the multipath command once and cache the output.
    multipathOutput, err := executeCommand("multipath", "-ll")
    if err != nil {
        return err
    }

    // Loop through the volume groups.
    volumeGroups, err := getVolumeGroups()
    if err != nil {
        return err
    }

    for _, vg := range volumeGroups {
        // Find the PVs in the VG
        pvOutput, err := executeCommand("vgdisplay", "-v", vg)
        if err != nil {
            return err
        }

        pvRegex := regexp.MustCompile(`PV Name\s+(\S+)`)
        pvMatches := pvRegex.FindAllStringSubmatch(pvOutput, -1)

        for _, match := range pvMatches {
            if len(match) < 2 {
                continue
            }

            pv := match[1]
            pvCleaned := strings.Split(pv, "/")
            pvName := pvCleaned[len(pvCleaned)-1]

            // Clean up the PV name by dropping partition numbers.
            pvName = regexp.MustCompile(`p[0-9]$`).ReplaceAllString(pvName, "")

            // Extract the WWID from the multipath output.
            wwidRegex := regexp.MustCompile(fmt.Sprintf(`%s\s+dm-(\S+)`, pvName))
            wwidMatches := wwidRegex.FindStringSubmatch(multipathOutput)
            if len(wwidMatches) < 2 {
                continue
            }
            wwid := wwidMatches[1]

            lastFourDigits := wwid[len(wwid)-4:]
            lastFiveDigits := wwid[len(wwid)-5:]

            // Only act on those not configured in multipathConfFile
            confFileContent, err := os.ReadFile(multipathConfFile)
            if err != nil {
                return err
            }
            if strings.Contains(string(confFileContent), wwid) {
                continue
            }

            alias := fmt.Sprintf("%s_%s", vg, lastFourDigits)
            if strings.Contains(string(confFileContent), alias) {
                alias = fmt.Sprintf("%s_%s", vg, lastFiveDigits)
            }

            // Output the multipath configuration stanza
            fmt.Printf("  multipath {\n")
            fmt.Printf("    wwid \"%s\"\n", wwid)
            fmt.Printf("    alias %s\n", alias)
            fmt.Printf("  }\n")
        }
    }
    return nil
}

func main() {
    // Run as root check
    if os.Geteuid() != 0 {
        fmt.Println("Please run with sudo or as root.")
        os.Exit(1)
    }

    multipathConfFile := "/etc/multipath.conf"

    // Generate multipath aliases
    if err := generateMultipathAliases(multipathConfFile); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

