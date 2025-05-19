package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sshdctl {start|stop}")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "start":
		startSSHD()
	case "stop":
		stopSSHD()
	default:
		fmt.Println("Invalid command. Use 'start' or 'stop'.")
		os.Exit(1)
	}
}

func startSSHD() {
	cmd := exec.Command("/usr/sbin/sshd")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error starting sshd: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("sshd started.")
}

func stopSSHD() {
	pid, err := readPID()
	if err != nil {
		fmt.Printf("Error reading PID: %v\n", err)
		os.Exit(1)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("Error finding process: %v\n", err)
		os.Exit(1)
	}

	err = process.Signal(syscall.SIGINT) // Use SIGTERM or SIGINT to gracefully stop the service
	if err != nil {
		fmt.Printf("Error stopping sshd: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("sshd stopped.")
}

func readPID() (int, error) {
	data, err := ioutil.ReadFile("/var/run/sshd.pid")
	if err != nil {
		return 0, err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}
	return pid, nil
}

