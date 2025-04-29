package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//This function takes an error as input and checks if it is of type *exec.ExitError.
//If it is, it extracts and returns the ProcessState from the error. If not, it returns nil.
func processState(e error) *os.ProcessState {
	err, ok := e.(*exec.ExitError)
	if !ok {
		return nil
	}
	return err.ProcessState
}

//This function takes a *os.ProcessState and checks if it can be asserted to syscall.WaitStatus.
//If it can, it extracts and returns the exit status of the process.
//If it cannot, it returns -1.

func exitStatus(state *os.ProcessState) int {
	status, ok := state.Sys().(syscall.WaitStatus)
	if !ok {
		return -1
	}
	return status.ExitStatus()
}

func main() {
	cmd := exec.Command("ls", "__a__")
	if err := cmd.Run(); err != nil {
		state := processState(err)
		if state == nil {
			fmt.Println(err)
			return
		}
		if status := exitStatus(state); status == -1 {
			fmt.Println(err)
		} else {
			fmt.Println("Status:", status)
		}
	}
}

// exec.Command("ls", "__a__"): Creates a command to list the contents of a directory named __a__, which
// likely does not exist.
// cmd.Run(): Runs the command and waits for it to complete.
// If an error occurs (likely because the directory __a__ does not exist), the error is passed to 
//processState to extract the process state.
// The process state is then passed to exitStatus to get the exit status code.
// Depending on the result, it prints either the error message or the exit status code.


// The status code 2 indicates that the ls command failed because the directory __a__ does not exist. 
// In Unix-like operating systems, the ls command 
// returns certain exit codes that help determine the nature of the error:

// 0: Success
// 1: Minor problems (e.g., no files matched the specified pattern)
// 2: Serious trouble (e.g., specified directory does not exist)
// In your Go program, the exec.Command("ls", "__a__") command is trying to list the contents 
// of a directory named __a__. Since this directory does not exist, the ls command fails and returns 
// an exit status of 2.