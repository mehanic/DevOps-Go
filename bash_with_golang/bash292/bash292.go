package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// First command: ls /usr
	cmd := "ls /usr"
	fmt.Printf("Output of command %s -\n", cmd)
	runCommand(cmd)

	// Second command: ls /usr | wc -l
	cmd1 := "ls /usr | wc -l"
	fmt.Printf("Line count of /usr -\n")
	runCommandWithShell(cmd1)

	// Expression: expr 2 + 4 * 6
	expression := "expr 2 + 4 \\* 6"
	fmt.Printf("Value of %s\n", expression)
	runCommandWithShell(expression)
}

// runCommand executes a simple command without pipes or shell interpretation.
func runCommand(command string) {
	// Split the command into name and arguments
	cmd := exec.Command("sh", "-c", command) // Use 'sh -c' to execute the string command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", err)
		return
	}
	fmt.Println(string(output))
}

// runCommandWithShell runs more complex commands that may involve pipes, like 'ls | wc -l'
func runCommandWithShell(command string) {
	// 'sh -c' allows executing the command through a shell
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing command: %s\n", err)
		return
	}
	fmt.Println(string(output))
}

