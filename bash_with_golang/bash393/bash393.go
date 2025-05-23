package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
//	"strings"
)

func main() {
	// Let's run a command and send all of the output to /dev/null
	fmt.Println("No output?")
	// Run the command to check for a fake file
	err := exec.Command("ls", "~/fakefile.txt").Run()
	if err != nil {
		// Redirecting error to /dev/null, just suppress output
		// In Go, we typically handle errors with logging or user feedback
		_ = exec.Command("true").Run() // This is a placeholder to suppress the output
	}

	// Retrieve output from a piped command
	fmt.Println("part 1")
	historyText, err := exec.Command("bash", "-c", "cat ~/.bashrc | grep HIST").Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	historyOutput := string(historyText)
	fmt.Println(historyOutput)

	// Output the results to history.config
	fmt.Println("part 2")
	err = ioutil.WriteFile("history.config", []byte(historyOutput), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Re-direct history.config as input to the cat command
	content, err := ioutil.ReadFile("history.config")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content))

	// Append a string to history.config
	err = appendToFile("history.config", "MY_VAR=1\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}

	fmt.Println("part 3 - using Tee")
	// Neato.txt will contain the same information as the console
	neatoFile, err := os.Create("neato.txt")
	if err != nil {
		fmt.Println("Error creating neato.txt:", err)
		return
	}
	defer neatoFile.Close()

	// Run the command and redirect output to neato.txt
	cmd := exec.Command("ls", "-la", "~/fakefile.txt", "~")
	cmd.Stdout = neatoFile
	cmd.Stderr = os.Stderr // Also print stderr to console
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
}

// appendToFile appends a string to a specified file.
func appendToFile(filename string, text string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(text); err != nil {
		return err
	}

	return nil
}

