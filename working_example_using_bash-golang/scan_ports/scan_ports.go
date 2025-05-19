package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Compiled Regular Expressions
var (
	reportRe  = regexp.MustCompile(`Nmap scan report for (.*)`)
	targetRe  = regexp.MustCompile(`[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+[0-9/]*`)
	gnmapRe   = regexp.MustCompile(`Host: (.*)Ports:`)
	versionRe = regexp.MustCompile(`# Nmap 6.25 scan initiated`)
	hostRe    = regexp.MustCompile(`Host: (.*) .*Ports:`)
	portsRe   = regexp.MustCompile(`Ports: (.*)\sIgnored State:`)
	osRe      = regexp.MustCompile(`OS: (.*)\sSeq Index:`)
)

func runCommand(cmd string) (string, string, error) {
	fmt.Println(cmd)
	// Run the command and capture the output
	output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		return "", "", err
	}

	return "", string(output), nil
}

func printWarnings(warnings string) {
	for _, w := range strings.Split(warnings, "\n") {
		if w == "" {
			continue
		}
		fmt.Printf("[-] %s\n", w)
		if w == "QUITTING!" {
			os.Exit(1)
		}
	}
}

func saveTargets(fileName string, ips []string) {
	fmt.Printf("[*] Saving live targets to %s\n", fileName)
	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer outFile.Close()

	for _, ip := range ips {
		outFile.WriteString(ip + "\n")
	}
}

func parsePorts(portStr string, broken bool) []string {
	ports := []string{}
	for _, port := range strings.Split(portStr, ",") {
		var num, stat, proto, sn, serv string
		parts := strings.Split(port, "/")

		if broken {
			num, stat, proto, sn, serv = parts[0], parts[1], parts[2], parts[3], parts[5]
		} else {
			num, stat, proto, sn, serv = parts[0], parts[1], parts[2], parts[4], parts[6]
		}

		service := serv
		if service == "" {
			service = sn
		}

		portEntry := fmt.Sprintf("%s/%s (%s) - %s", proto, strings.TrimSpace(num), stat, service)
		ports = append(ports, portEntry)
	}

	return ports
}

func parseGnmap(fileName string) (map[string]map[string]interface{}, error) {
	hosts := make(map[string]map[string]interface{})
	broken := false

	file, err := os.Open(fileName + ".gnmap")
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
			h := hostRe.FindStringSubmatch(line)
			host := "Unknown"
			if len(h) > 1 {
				host = h[1]
			}

			// Get Ports
			p := portsRe.FindStringSubmatch(line)
			var ports []string
			if p != nil {
				ports = parsePorts(p[1], broken)
			}

			// Get OS
			o := osRe.FindStringSubmatch(line)
			os := "Unknown"
			if len(o) > 1 {
				os = o[1]
			}

			hosts[host] = map[string]interface{}{
				"os":    os,
				"ports": ports,
			}
		}
	}

	return hosts, nil
}

func main() {
	// Command-line argument parsing
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: discover IP_addresses <port_list>")
		return
	}

	target := os.Args[1]
	var otherPorts string
	if len(os.Args) == 3 {
		otherPorts = os.Args[2]
	}

	pingFileName := strings.ReplaceAll(fmt.Sprintf("%s_ping_scan", target), "/", ".")
	targetFileName := strings.ReplaceAll(fmt.Sprintf("%s_targets.txt", target), "/", ".")
	synFileName := strings.ReplaceAll(fmt.Sprintf("%s_syn_scan", target), "/", ".")
	resultFileName := strings.ReplaceAll(fmt.Sprintf("%s_results.md", target), "/", ".")

	// Default ports
	ports := "21,22,23,25,53,80,110,119,143,443,135,139,445,593,1352,1433,1498," +
		"1521,3306,5432,389,1494,1723,2049,2598,3389,5631,5800,5900,6000"
	if otherPorts != "" {
		ports += "," + otherPorts
	}

	fmt.Printf("[*] Running discovery scan against targets %s\n", target)

	// Check if the target is a valid IP range
	if targetRe.MatchString(target) {
		cmd := fmt.Sprintf("nmap -sn -PE -n -oA %s %s", pingFileName, target)
		_, resp, err := runCommand(cmd)
		if err != nil {
			fmt.Println("Error running command:", err)
			return
		}
		printWarnings(resp)

		ips := reportRe.FindAllStringSubmatch(resp, -1)
		fmt.Printf("[+] Found %d live targets\n", len(ips))

		if len(ips) == 0 {
			fmt.Println("[-] No targets to scan. Quitting.")
			return
		}

		var liveIPs []string
		for _, ip := range ips {
			liveIPs = append(liveIPs, ip[1])
		}
		saveTargets(targetFileName, liveIPs)

		fmt.Println("[*] Ping scan complete.\n")
	} else {
		cmd := fmt.Sprintf("nmap -sn -PE -n -A %s -iL %s", pingFileName, target)
		_, resp, err := runCommand(cmd)
		if err != nil {
			fmt.Println("Error running command:", err)
			return
		}
		printWarnings(resp)
	}

	fmt.Printf("[*] Running full scan on live addresses using ports %s\n", ports)
	cmd := fmt.Sprintf("nmap -sS -n -A -p %s --open -oA %s -iL %s", ports, synFileName, targetFileName)
	_, resp, err := runCommand(cmd)
	if err != nil {
		fmt.Println("Error running command:", err)
		return
	}
	printWarnings(resp)
	fmt.Println("[*] Full scan complete.\n")

	fmt.Println("[*] Parsing Scan results.")
	hosts, err := parseGnmap(synFileName)
	if err != nil {
		fmt.Println("Error parsing gnmap:", err)
		return
	}

	fmt.Printf("[*] Saving results to %s\n", resultFileName)
	outFile, err := os.Create(resultFileName)
	if err != nil {
		fmt.Println("Error creating results file:", err)
		return
	}
	defer outFile.Close()

	for host, data := range hosts {
		outFile.WriteString(host + "\n")
		outFile.WriteString(strings.Repeat("=", len(host)) + "\n\n")
		outFile.WriteString("OS\n")
		outFile.WriteString("--\n")
		outFile.WriteString(data["os"].(string) + "\n\n")
		outFile.WriteString("Ports\n")
		outFile.WriteString("-----\n")
		for _, port := range data["ports"].([]string) {
			outFile.WriteString(port + "\n")
		}
		outFile.WriteString("\n\n")
	}

	fmt.Println("[*] Parsing results is complete.")
}

