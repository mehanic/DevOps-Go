package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
)

func main() {
	remoteHost := "host.example.com"  // Required
	remoteFile := "/etc/passwd"        // Required
	sshUser := "user@"                  // Optional, set to "" to not use
	sshID := ""                        // Optional, set to "" to not use

	// Create a new SSH client configuration
	var config *ssh.ClientConfig
	if sshID != "" {
		key, err := ioutil.ReadFile(sshID)
		if err != nil {
			fmt.Printf("Unable to read SSH key: %v\n", err)
			os.Exit(1)
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			fmt.Printf("Unable to parse SSH key: %v\n", err)
			os.Exit(1)
		}

		config = &ssh.ClientConfig{
			User: sshUser,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Note: insecure; use a proper verification in production!
		}
	} else {
		config = &ssh.ClientConfig{
			User: sshUser,
			Auth: []ssh.AuthMethod{
				ssh.Password("your-password-here"), // Replace with your password or manage via SSH agent
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Note: insecure; use a proper verification in production!
		}
	}

	// Connect to the remote host
	client, err := ssh.Dial("tcp", remoteHost+":22", config)
	if err != nil {
		fmt.Printf("Failed to dial SSH: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	// Create a new session
	session, err := client.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %v\n", err)
		os.Exit(1)
	}
	defer session.Close()

	// Command to check if the file is readable
	cmd := fmt.Sprintf("[ -r %s ] && echo 1 || echo 0", remoteFile)
	result, err := session.Output(cmd)
	if err != nil {
		fmt.Printf("SSH command failed: %v\n", err)
		os.Exit(1)
	}

	// Check the result
	if string(result) == "1\n" {
		fmt.Printf("%s present on %s\n", remoteFile, remoteHost)
	} else {
		fmt.Printf("%s not present on %s\n", remoteFile, remoteHost)
	}
}



//#!/usr/bin/env bash
//# cookbook filename: command_substitution

//REMOTE_HOST='host.example.com'  # Required
//REMOTE_FILE='/etc/passwd'       # Required
//SSH_USER='user@'                # Optional, set to '' to not use
//#SSH_ID='-i ~/.ssh/foo.id'       # Optional, set to '' to not use
//SSH_ID=''

//result=$(
//    ssh $SSH_ID $SSH_USER$REMOTE_HOST \
//      "[ -r $REMOTE_FILE ] && echo 1 || echo 0"
//) || { echo "SSH command failed!" >&2; exit 1; }

//if [ $result = 1 ]; then
//    echo "$REMOTE_FILE present on $REMOTE_HOST"
//else
//    echo "$REMOTE_FILE not present on $REMOTE_HOST"
//fi
