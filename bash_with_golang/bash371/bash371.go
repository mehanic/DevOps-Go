package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Constants
const (
	hostnamePrefix = "System Information for "
)

// Function to get the system release information
func systemInfo() string {
	return "<h2>System release info</h2><p>Function not yet implemented</p>"
}

// Function to get the system uptime
func showUptime() string {
	uptimeCmd := exec.Command("uptime")
	uptimeOutput, err := uptimeCmd.Output()
	if err != nil {
		return "<h2>System uptime</h2><p>Error retrieving uptime.</p>"
	}
	return "<h2>System uptime</h2><pre>" + string(uptimeOutput) + "</pre>"
}

// Function to get drive space
func driveSpace() string {
	dfCmd := exec.Command("df")
	dfOutput, err := dfCmd.Output()
	if err != nil {
		return "<h2>Filesystem space</h2><p>Error retrieving filesystem space.</p>"
	}
	return "<h2>Filesystem space</h2><pre>" + string(dfOutput) + "</pre>"
}

// Function to get home directory space by user
func homeSpace() string {
	if os.Geteuid() != 0 { // Check if the user is root
		return ""
	}
	cmd := exec.Command("du", "-s", "/home/*")
	duOutput, err := cmd.Output()
	if err != nil {
		return "<h2>Home directory space by user</h2><p>Error retrieving directory space.</p>"
	}

	lines := strings.Split(string(duOutput), "\n")
	spaceMap := make(map[int]string)
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) > 1 {
			if bytes, err := strconv.Atoi(parts[0]); err == nil {
				spaceMap[bytes] = parts[1]
			}
		}
	}

	// Sort the space map
	keys := make([]int, 0, len(spaceMap))
	for k := range spaceMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var sb strings.Builder
	sb.WriteString("<h2>Home directory space by user</h2><pre>")
	sb.WriteString("Bytes Directory\n")
	for _, key := range keys {
		sb.WriteString(fmt.Sprintf("%d %s\n", key, spaceMap[key]))
	}
	sb.WriteString("</pre>")
	return sb.String()
}

// Main function to generate HTML
func main() {
	hostname, err := os.Hostname() // Get the hostname and handle error
	if err != nil {
		hostname = "Unknown Host"
	}
	title := hostnamePrefix + hostname
	rightNow := time.Now().Format("02-Jan-2006 15:04:05 MST")
	timeStamp := "Updated on " + rightNow + " by " + os.Getenv("USER")

	fmt.Println("<html>")
	fmt.Println("<head>")
	fmt.Println("<title>" + title + "</title>")
	fmt.Println("</head>")
	fmt.Println("<body>")
	fmt.Println("<h1>" + title + "</h1>")
	fmt.Println("<p>" + timeStamp + "</p>")
	fmt.Println(systemInfo())
	fmt.Println(showUptime())
	fmt.Println(driveSpace())
	fmt.Println(homeSpace())
	fmt.Println("</body>")
	fmt.Println("</html>")
}

