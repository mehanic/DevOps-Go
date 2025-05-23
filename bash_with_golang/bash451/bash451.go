package main

import (
	"fmt"
	"os/exec"
//	"runtime"
)

// execCommand runs a command and handles errors
func execCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing command '%s %s': %v\nOutput: %s", command, args, err, output)
	}
	return nil
}

func main() {
	// Set default policy to DROP
	fmt.Println("Setting default policies to DROP...")
	if err := execCommand("iptables", "-P", "INPUT", "DROP"); err != nil {
		fmt.Println(err)
		return
	}
	if err := execCommand("iptables", "-P", "OUTPUT", "DROP"); err != nil {
		fmt.Println(err)
		return
	}
	if err := execCommand("iptables", "-P", "FORWARD", "DROP"); err != nil {
		fmt.Println(err)
		return
	}

	// Enable IP forwarding
	fmt.Println("Enabling IP forwarding...")
	if err := execCommand("sysctl", "-w", "net.ipv4.ip_forward=1"); err != nil {
		fmt.Println(err)
		return
	}

	// Allow traffic from internal (eth0) to DMZ (eth2)
	fmt.Println("Allowing traffic from eth0 to eth2...")
	if err := execCommand("iptables", "-t", "filter", "-A", "FORWARD", "-i", "eth0", "-o", "eth2", "-m", "state", "--state", "NEW,ESTABLISHED,RELATED", "-j", "ACCEPT"); err != nil {
		fmt.Println(err)
		return
	}
	if err := execCommand("iptables", "-t", "filter", "-A", "FORWARD", "-i", "eth2", "-o", "eth0", "-m", "state", "--state", "ESTABLISHED,RELATED", "-j", "ACCEPT"); err != nil {
		fmt.Println(err)
		return
	}

	// Allow traffic from internet (ens33) to DMZ (eth2)
	fmt.Println("Allowing traffic from ens33 to eth2...")
	if err := execCommand("iptables", "-t", "filter", "-A", "FORWARD", "-i", "ens33", "-o", "eth2", "-m", "state", "--state", "NEW,ESTABLISHED,RELATED", "-j", "ACCEPT"); err != nil {
		fmt.Println(err)
		return
	}
	if err := execCommand("iptables", "-t", "filter", "-A", "FORWARD", "-i", "eth2", "-o", "ens33", "-m", "state", "--state", "ESTABLISHED,RELATED", "-j", "ACCEPT"); err != nil {
		fmt.Println(err)
		return
	}

	// Redirect incoming web requests from ens33 to web server
	fmt.Println("Redirecting web requests...")
	if err := execCommand("iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "80", "-j", "DNAT", "--to-dest", "192.168.20.2"); err != nil {
		fmt.Println(err)
		return
	}
	if err := execCommand("iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "443", "-j", "DNAT", "--to-dest", "192.168.20.2"); err != nil {
		fmt.Println(err)
		return
	}

	// Redirect incoming mail (SMTP) requests from ens33 to Mail server
	fmt.Println("Redirecting mail requests...")
	if err := execCommand("iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "25", "-j", "DNAT", "--to-dest", "192.168.20.3"); err != nil {
		fmt.Println(err)
		return
	}

	// Redirect incoming DNS requests from ens33 to DNS server
	fmt.Println("Redirecting DNS requests...")
	if err := execCommand("iptables", "-t", "nat", "-A", "PREROUTING", "-p", "udp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "53", "-j", "DNAT", "--to-dest", "192.168.20.4"); err != nil {
		fmt.Println(err)
		return
	}
	if err := execCommand("iptables", "-t", "nat", "-A", "PREROUTING", "-p", "tcp", "-i", "ens33", "-d", "200.0.0.1", "--dport", "53", "-j", "DNAT", "--to-dest", "192.168.20.4"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Firewall rules have been successfully applied.")
}

