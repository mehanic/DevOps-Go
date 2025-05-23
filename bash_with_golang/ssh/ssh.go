package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Check if there is one argument (the SSH server address)
	if len(os.Args) != 2 {
		fmt.Println("対象のsshサーバーのIPアドレスをパラメータで付与して下さい")
		os.Exit(1)
	}

	sshServer := os.Args[1]

	// Step 1: Generate SSH key (DSA)
	fmt.Println("Generating SSH key...")
	cmd := exec.Command("ssh-keygen", "-t", "dsa", "-f", os.Getenv("HOME")+"/.ssh/id_dsa", "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error generating SSH key: %v", err)
	}

	// Step 2: Prepare remote .ssh directory and set correct permissions
	fmt.Println("Preparing remote .ssh directory...")
	cmd = exec.Command("ssh", sshServer, `if ! [ -d ~/.ssh ]; then mkdir .ssh ; fi ; chmod 0700 .ssh`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error preparing remote .ssh directory: %v", err)
	}

	// Step 3: Copy the public key to the remote server
	fmt.Println("Copying public key to remote server...")
	cmd = exec.Command("scp", os.Getenv("HOME")+"/.ssh/id_dsa.pub", sshServer+":~/.ssh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error copying public key to remote server: %v", err)
	}

	// Step 4: Add the public key to authorized_keys and set correct permissions
	fmt.Println("Configuring authorized_keys on remote server...")
	cmd = exec.Command("ssh", sshServer, `cat ~/.ssh/id_dsa.pub >> ~/.ssh/authorized_keys ; rm -f ~/.ssh/id_dsa.pub ; chmod 0600 ~/.ssh/authorized_keys`)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error configuring authorized_keys on remote server: %v", err)
	}

	// Step 5: Set local private key permissions to 0600
	fmt.Println("Setting local private key permissions...")
	cmd = exec.Command("chmod", "0600", os.Getenv("HOME")+"/.ssh/id_dsa")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error setting local private key permissions: %v", err)
	}

	fmt.Println("SSH key setup complete.")
}

