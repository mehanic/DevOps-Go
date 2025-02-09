package main

import (
	"fmt"
	"os"
	"os/exec"
	// "strings"
)

func main() {
	// Let's run a command and send all of the output to /dev/null
	fmt.Println("No output?")
	err := exec.Command("ls", "/home/username/fakefile.txt").Run() // Replace "username" with your actual username
	if err != nil {
		// Suppressing the output by ignoring the error
	}

	// Retrieve output from a piped command
	fmt.Println("part 1")
	historyText, err := exec.Command("bash", "-c", "cat ~/.bashrc | grep HIST").Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(historyText))

	// Output the results to history.config
	fmt.Println("part 2")
	if err = os.WriteFile("history.config", historyText, 0644); err != nil {
		fmt.Println("Error writing to history.config:", err)
		return
	}

	// Re-direct history.config as input to the cat command
	content, err := os.ReadFile("history.config")
	if err != nil {
		fmt.Println("Error reading history.config:", err)
		return
	}
	fmt.Println(string(content))

	// Append a string to history.config
	err = os.WriteFile("history.config", append(content, []byte("MY_VAR=1\n")...), 0644)
	if err != nil {
		fmt.Println("Error appending to history.config:", err)
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

	lsCommand := exec.Command("ls", "-la", "/home/username/fakefile.txt", "/home/username") // Replace "username" with your actual username
	lsOutput, err := lsCommand.CombinedOutput()
	if err != nil {
		fmt.Println("Error running ls command:", err)
		return
	}

	// Print to console and write to neato.txt
	fmt.Print(string(lsOutput))
	neatoFile.Write(lsOutput)
}
