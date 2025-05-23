package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("iptables", "-A", "INPUT", "-p", "tcp", "--dport", "80", "-j", "ACCEPT")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

