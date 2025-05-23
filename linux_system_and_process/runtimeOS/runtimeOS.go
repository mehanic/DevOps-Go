package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
//	"time"
)

func mycode(cmd1, cmd2 string) {
	fmt.Println("Please wait.")
//	time.Sleep(2 * time.Second)

	// Execute the first command
	cmd := exec.Command(cmd1)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println("Please wait finding the list of dir and files")
	//	time.Sleep(2 * time.Second)

	// Execute the second command
	cmd = exec.Command(cmd2)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	if runtime.GOOS == "windows" {
		mycode("cmd", "/c cls && dir")
	} else {
		mycode("ls", "-la")
	}
}

