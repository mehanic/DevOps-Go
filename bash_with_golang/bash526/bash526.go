package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
//	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("対象のsshサーバーのIPアドレスをパラメータで付与して下さい")
		return
	}

	sshServer := os.Args[1]

	// Step 1: Generate DSA key
	if err := exec.Command("ssh-keygen", "-t", "dsa", "-N", "", "-f", filepath.Join(os.Getenv("HOME"), ".ssh", "id_dsa")).Run(); err != nil {
		fmt.Println("Error generating SSH key:", err)
		return
	}

	// Step 2: Ensure .ssh directory exists on remote server
	if err := exec.Command("ssh", sshServer, "mkdir -p ~/.ssh && chmod 0700 ~/.ssh").Run(); err != nil {
		fmt.Println("Error creating .ssh directory on server:", err)
		return
	}

	// Step 3: Copy public key to the remote server
	if err := exec.Command("scp", filepath.Join(os.Getenv("HOME"), ".ssh", "id_dsa.pub"), fmt.Sprintf("%s:~/.ssh/", sshServer)).Run(); err != nil {
		fmt.Println("Error copying public key to server:", err)
		return
	}

	// Step 4: Append public key to authorized_keys and set permissions
	if err := exec.Command("ssh", sshServer, "cat ~/.ssh/id_dsa.pub >> ~/.ssh/authorized_keys && rm -f ~/.ssh/id_dsa.pub && chmod 0600 ~/.ssh/authorized_keys").Run(); err != nil {
		fmt.Println("Error setting up authorized_keys on server:", err)
		return
	}

	// Step 5: Set local key permissions
	if err := os.Chmod(filepath.Join(os.Getenv("HOME"), ".ssh", "id_dsa"), 0600); err != nil {
		fmt.Println("Error setting local key permissions:", err)
		return
	}

	fmt.Println("SSH key setup completed successfully.")
}

