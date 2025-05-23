// createchildprocess.go
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("/bin/ls", "-l")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Process ID: %d, Group ID: %d\n", cmd.Process.Pid, -cmd.Process.Pid)
}

