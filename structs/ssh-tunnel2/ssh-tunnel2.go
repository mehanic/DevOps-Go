package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var config = struct {
	SSHServer  string
	SSHPort    int
	SSHUser    string
	SSHKey     string
	RemotePort int
	UnixSocket string
}{
	SSHServer:  "ec2-18-196-140-221.eu-central-1.compute.amazonaws.com",
	SSHPort:    22,
	SSHUser:    "ec2-user",
	SSHKey:     "q.pem",
	RemotePort: 62222,
	UnixSocket: "/tmp/ssh_tunnel.sock",
}

func main() {
	if err := start(); err != nil {
		log.Fatalf("Failed to start SSH tunnel: %v", err)
	}
	fmt.Println("SSH tunnel started.")
}

func start() error {
	// Open /dev/null for discarding output
	devNull, err := os.OpenFile(os.DevNull, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open /dev/null: %v", err)
	}
	defer devNull.Close() // Ensure it's closed when done

	cmd := exec.Command("ssh",
		config.SSHServer,
		"-N",
		"-i", config.SSHKey,
		"-M", "-S", config.UnixSocket,
		"-o", "UserKnownHostsFile=/dev/null",
		"-o", "StrictHostKeyChecking=no",
		"-o", "ExitOnForwardFailure=yes",
		"-p", fmt.Sprint(config.SSHPort),
		"-l", config.SSHUser,
		"-R", fmt.Sprintf("%d:localhost:22", config.RemotePort),
	)

	cmd.Stdout = devNull
	cmd.Stderr = devNull

	return cmd.Start()
}

func stop() error {
	return controlSSH("exit")
}

func status() error {
	return controlSSH("check")
}

func controlSSH(command string) error {
	// Open /dev/null for discarding output
	devNull, err := os.OpenFile(os.DevNull, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open /dev/null: %v", err)
	}
	defer devNull.Close() // Ensure it's closed when done

	cmd := exec.Command("ssh", "-S", config.UnixSocket, "-O", command, "x")
	cmd.Stdout = devNull
	cmd.Stderr = devNull

	return cmd.Run()
}

