package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// Function to run a shell command and return the output
func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func main() {
	// Display the date
	currentTime := time.Now()
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))

	// Display uptime
	uptime, _ := runCommand("uptime")
	fmt.Println("uptime:")
	fmt.Println(uptime)

	// Display currently connected users
	w, _ := runCommand("w")
	fmt.Println("Currently connected:")
	fmt.Println(w)

	fmt.Println("--------------------")
	
	// Display last logins
	lastLogins, _ := runCommand("last", "-a")
	lastLoginsLines := strings.Split(lastLogins, "\n")
	for i, line := range lastLoginsLines {
		if i >= 3 { // Get the first three lines
			break
		}
		fmt.Println(line)
	}

	fmt.Println("--------------------")

	// Display disk and memory usage
	df, _ := runCommand("df", "-h")
	dfLines := strings.Fields(df)
	fmt.Printf("Free/total disk: %s / %s\n", dfLines[len(dfLines)-11], dfLines[len(dfLines)-9]) // Adjust indices as necessary

	free, _ := runCommand("free", "-m")
	freeLines := strings.Fields(free)
	fmt.Printf("Free/total memory: %s / %s MB\n", freeLines[16], freeLines[7]) // Adjust indices as necessary

	fmt.Println("--------------------")

	// OOM errors
	startLog, err := runCommand("head", "-1", "/var/log/messages")
	if err != nil {
		fmt.Println("Error reading /var/log/messages:", err)
		return
	}

	startLog = strings.TrimSpace(startLog)

	// Check length before slicing
	var startLogOutput string
	if len(startLog) >= 12 {
		startLogOutput = startLog[:12]
	} else {
		startLogOutput = startLog // Or handle this case differently
	}

	oomOutput, _ := runCommand("grep", "-ci", "kill", "/var/log/messages")
	oom := strings.TrimSpace(oomOutput)
	fmt.Printf("OOM errors since %s : %s\n", startLogOutput, oom)

	fmt.Println("--------------------")

	// Utilization and most expensive processes
	top, _ := runCommand("top", "-b")
	topLines := strings.Split(top, "\n")
	for i, line := range topLines {
		if i < 3 { // Get the first three lines
			fmt.Println(line)
		} else if i >= 6 && i < 10 { // Get lines 7-10 for most expensive processes
			fmt.Println(line)
		}
	}

	fmt.Println("--------------------")

	// Open TCP ports
	nmap, _ := runCommand("nmap", "-p-", "-T4", "127.0.0.1")
	fmt.Println("Open TCP ports:")
	fmt.Println(nmap)

	fmt.Println("--------------------")

	// Current connections
	ss, _ := runCommand("ss", "-s")
	fmt.Println("Current connections:")
	fmt.Println(ss)

	fmt.Println("--------------------")

	// Processes
	ps, _ := runCommand("ps", "auxf")
	fmt.Println("Processes:")
	fmt.Println(ps)

	fmt.Println("--------------------")

	// VMStat
	vmstat, _ := runCommand("vmstat", "1", "5")
	fmt.Println("vmstat:")
	fmt.Println(vmstat)
}

