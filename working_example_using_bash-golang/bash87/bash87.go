package main

import (
	"log"
	"os/exec"
)

func main() {
	// Set the default policy to DROP
	commands := [][]string{
		{"iptables", "-P", "INPUT", "DROP"},
		{"iptables", "-P", "OUTPUT", "DROP"},
		{"iptables", "-P", "FORWARD", "DROP"},
		// Enable IP forwarding
		{"sysctl", "-w", "net.ipv4.ip_forward=1"},
		// Allow traffic from internal (eth0) to DMZ (eth2)
		{"iptables", "-t", "filter", "-A", "FORWARD", "-i", "eth0", "-o", "eth2", "-m", "state", "--state", "NEW,ESTABLISHED,RELATED", "-j", "ACCEPT"},
		{"iptables", "-t", "filter", "-A", "FORWARD", "-i", "eth2", "-o", "eth0", "-m", "state", "--state", "ESTABLISHED,RELATED", "-j", "ACCEPT"},
		// Allow traffic from internet (ens33) to DMZ (eth2)
		{"iptables", "-t", "filter", "-A", "FORWARD", "-i", "ens33", "-o", "eth2", "-m", "state", "--state", "NEW,ESTABLISHED,RELATED", "-j", "ACCEPT"},
		{"iptables", "-t", "filter", "-A", "FORWARD", "-i", "eth2", "-o", "ens33", "-m", "state", "--state", "ESTABLISHED,RELATED", "-j", "ACCEPT"},
		// Redirect incoming web requests to web server
		{"iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "80", "-j", "DNAT", "--to-dest", "192.168.20.2"},
		{"iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "443", "-j", "DNAT", "--to-dest", "192.168.20.2"},
		// Redirect incoming mail requests to Mail server
		{"iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "25", "-j", "DNAT", "--to-dest", "192.168.20.3"},
		// Redirect incoming DNS requests to DNS server
		{"iptables", "-t", "nat", "-A", "PREROUTING", "-p", "udp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "53", "-j", "DNAT", "--to-dest", "192.168.20.4"},
		{"iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "53", "-j", "DNAT", "--to-dest", "192.168.20.4"},
	}

	for _, cmd := range commands {
		if err := executeCommand(cmd); err != nil {
			log.Fatalf("Error executing command %v: %v", cmd, err)
		}
	}
}

// executeCommand executes the given command and logs any error.
func executeCommand(cmd []string) error {
	command := exec.Command(cmd[0], cmd[1:]...) // Create command
	output, err := command.CombinedOutput()      // Execute command and get combined output
	if err != nil {
		return err
	}
	log.Printf("Executed: %s, Output: %s", cmd, output)
	return nil
}

