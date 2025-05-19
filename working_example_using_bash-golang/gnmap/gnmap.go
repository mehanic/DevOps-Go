package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//-----------------------------------------------------------------------------
// Compiled Regular Expressions
//-----------------------------------------------------------------------------
var (
	gnmapRe   = regexp.MustCompile(`Host: (.*)Ports:`)
	versionRe = regexp.MustCompile(`# Nmap 6.25 scan initiated`)
	hostRe    = regexp.MustCompile(`Host: (.*) .*Ports:`)
	portsRe   = regexp.MustCompile(`Ports: (.*)\sIgnored State:`)
	osRe      = regexp.MustCompile(`OS: (.*)\sSeq Index:`)
)

//-----------------------------------------------------------------------------
// Functions
//-----------------------------------------------------------------------------

// parsePorts processes the ports string and handles the differences in the broken format.
func parsePorts(portStr string, broken bool) []string {
	ports := []string{}
	portList := strings.Split(portStr, ",")
	for _, port := range portList {
		var service string

		if broken {
			fields := strings.Split(port, "/")
			if len(fields) >= 6 {
				// We're only using these fields as an example.
				num, stat, proto, _, sn, serv := fields[0], fields[1], fields[2], fields[3], fields[4], fields[5]
				service = getService(serv, sn)
				portStr := fmt.Sprintf("%s/%s (%s) - %s", proto, strings.TrimSpace(num), stat, service)
				ports = append(ports, portStr)
			}
		} else {
			fields := strings.Split(port, "/")
			if len(fields) >= 8 {
				num, stat, proto, _, sn, _, serv := fields[0], fields[1], fields[2], fields[3], fields[4], fields[6], fields[7]
				service = getService(serv, sn)
				portStr := fmt.Sprintf("%s/%s (%s) - %s", proto, strings.TrimSpace(num), stat, service)
				ports = append(ports, portStr)
			}
		}
	}
	return ports
}

// getService returns the correct service field value.
func getService(serv string, sn string) string {
	if serv == "" {
		return sn
	}
	return serv
}

// parseGnmap reads the gnmap file and parses the hosts, OS, and ports.
func parseGnmap(fileName string) (map[string]map[string]interface{}, error) {
	hosts := make(map[string]map[string]interface{})
	broken := false

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if versionRe.MatchString(line) {
			broken = true
		}

		if gnmapRe.MatchString(line) {
			// Get Hostname
			hostMatch := hostRe.FindStringSubmatch(line)
			host := "Unknown"
			if hostMatch != nil {
				host = hostMatch[1]
			}

			// Get Ports
			portsMatch := portsRe.FindStringSubmatch(line)
			ports := []string{}
			if portsMatch != nil {
				ports = parsePorts(portsMatch[1], broken)
			}

			// Get OS
			osMatch := osRe.FindStringSubmatch(line)
			os := "Unknown"
			if osMatch != nil {
				os = osMatch[1]
			}

			hosts[host] = map[string]interface{}{
				"os":    os,
				"ports": ports,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return hosts, nil
}

//-----------------------------------------------------------------------------
// Main Program
//-----------------------------------------------------------------------------
func main() {
	usage := `
USAGE:

gnmap2md gnmap_file_name md_file_name

Converts a Nmap gnmap formatted file into a Markdown formatted file.
`
	if len(os.Args) != 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	gnmapFileName := os.Args[1]
	mdFileName := os.Args[2]

	fmt.Println("[*] Parsing Scan results.")
	hosts, err := parseGnmap(gnmapFileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("[*] Saving results to %s\n", mdFileName)
	outFile, err := os.Create(mdFileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for host, details := range hosts {
		// Write host
		writer.WriteString(host + "\n")
		writer.WriteString(strings.Repeat("=", len(host)) + "\n\n")

		// Write OS
		writer.WriteString("OS\n")
		writer.WriteString("--\n")
		writer.WriteString(details["os"].(string) + "\n\n")

		// Write Ports
		writer.WriteString("Ports\n")
		writer.WriteString("-----\n")
		for _, port := range details["ports"].([]string) {
			writer.WriteString(port + "\n")
		}
		writer.WriteString("\n\n")
	}

	writer.Flush()
	fmt.Println("[*] Parsing results is complete.")
}

