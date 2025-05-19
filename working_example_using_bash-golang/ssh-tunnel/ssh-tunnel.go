package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
)

func main() {
	// Define SSH server details
	sshHost := "ec2-18-196-140-221.eu-central-1.compute.amazonaws.com"
	sshUser := "ec2-user"
	privateKeyPath := "/home/petro/Downloads/GA/GuardianAngel.pem"

	// Load the private key
	key, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	// Set up the SSH client configuration
	config := &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // NOTE: Insecure, do not use in production!
	}

	// Establish SSH connection
	addr := fmt.Sprintf("%s:22", sshHost)
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatalf("failed to dial SSH: %v", err)
	}
	defer client.Close()

	// Set up port forwarding
	localListener, err := net.Listen("tcp", ":0") // Use an available local port
	if err != nil {
		log.Fatalf("failed to set up local listener: %v", err)
	}
	defer localListener.Close()

	localPort := localListener.Addr().(*net.TCPAddr).Port
	fmt.Printf("Local bind port: %d\n", localPort)

	// Forward traffic from the local port to the remote address
	go func() {
		for {
			// Accept incoming connections on the local port
			localConn, err := localListener.Accept()
			if err != nil {
				log.Fatalf("failed to accept local connection: %v", err)
			}

			// Connect to the remote address
			remoteConn, err := client.Dial("tcp", "127.0.0.1:22")
			if err != nil {
				log.Fatalf("failed to dial remote connection: %v", err)
			}

			// Handle the connection in a goroutine
			go func() {
				defer localConn.Close()
				defer remoteConn.Close()
				go func() {
					io.Copy(remoteConn, localConn) // Copy data from local to remote
				}()
				io.Copy(localConn, remoteConn) // Copy data from remote to local
			}()
		}
	}()

	// Keep the main function running
	select {}
}

