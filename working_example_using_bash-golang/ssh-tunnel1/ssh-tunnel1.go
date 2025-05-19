package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Define SSH server details
	hostname := "18.196.140.221"
	username := "ec2-user"
	privateKeyPath := ".pem"

	// Load the private key
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	// Set up the SSH client configuration
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // NOTE: Insecure, do not use in production!
	}

	// Establish SSH connection
	addr := fmt.Sprintf("%s:22", hostname)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatalf("failed to dial SSH: %v", err)
	}
	defer client.Close()

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}
	defer session.Close()

	// Execute the command
	output, err := session.CombinedOutput("python hello.py")
	if err != nil {
		log.Fatalf("failed to run command: %v", err)
	}

	// Print the output
	fmt.Println(string(output))
}
