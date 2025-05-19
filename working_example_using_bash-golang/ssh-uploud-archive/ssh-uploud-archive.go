package main

import (
//  "bufio"
	"fmt"
	"os"
	"os/exec"
//	"strings"
	"time"
)

// Global variables
var (
	scriptName   = "upload.go"
	description  = "Archive new or modified files and upload to web site"
	author       = "Chris F.A. Johnson"
	version      = "1.0"

	// Archive and upload settings
	host         = "127.0.0.1"                           // Remote host (URL or IP address)
	port         = "22"                                  // SSH port
	dest         = "work/upload"                         // Destination directory
	user         = "chris"                               // Login name on remote system
	source       = os.Getenv("HOME") + "/public_html/example.com" // Local directory to upload
	archiveDir   = os.Getenv("HOME") + "/work/webarchives"       // Directory to store archive files
	syncFile     = ".sync"                               // File to touch with time of last upload

	// Options
	menu         = false // do not print a menu
	qa           = false // do not use question and answer
	test         = false // 0 = upload for real; 1 = don't archive/upload, show settings
	configFile   string  // if defined, the file will be sourced
	configDir    = os.Getenv("HOME") + "/.config" // default location for configuration files
	sleepyTime   = 2     // delay in seconds after printing messages

	varInfo = map[string]string{
		"host":       "Remote host (URL or IP address)",
		"port":       "SSH port",
		"dest":       "Destination directory",
		"user":       "Login name on remote system",
		"source":     "Local directory to upload",
		"archiveDir": "Directory to store archive files",
		"syncFile":   "File to touch with time of last upload",
	}
)

func main() {
	// Parse command-line arguments
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-c":
			i++
			if i < len(os.Args) {
				configFile = os.Args[i]
			}
		case "-h":
			i++
			if i < len(os.Args) {
				host = os.Args[i]
			}
		case "-p":
			i++
			if i < len(os.Args) {
				port = os.Args[i]
			}
		case "-s":
			i++
			if i < len(os.Args) {
				source = os.Args[i]
			}
		case "-d":
			i++
			if i < len(os.Args) {
				dest = os.Args[i]
			}
		case "-u":
			i++
			if i < len(os.Args) {
				user = os.Args[i]
			}
		case "-a":
			i++
			if i < len(os.Args) {
				archiveDir = os.Args[i]
			}
		case "-f":
			i++
			if i < len(os.Args) {
				syncFile = os.Args[i]
			}
		case "-t":
			test = true
		case "-m":
			menu = true
		case "-q":
			qa = true
		}
	}

	// Load configuration if specified
	if configFile != "" {
		loadConfig(configFile)
	}

	// Execute menu or qa if defined
	if menu {
		showMenu()
	} else if qa {
		questionAnswer()
	}

	// Create datestamped filename for tarball
	tarFile := fmt.Sprintf("%s/%s.%s.tgz", archiveDir, host, time.Now().Format("2006-01-02T15:04:05"))

	if !test {
		if err := os.Chdir(source); err != nil {
			die(4, fmt.Sprintf("Failed to change directory to %s: %s", source, err))
		}
	}

	// Archive and upload files
	if !test {
		fmt.Printf("Archiving and uploading new files in directory: %s\n\n", source)
		if err := archiveAndUpload(tarFile); err != nil {
			die(3, err.Error())
		}
	} else {
		printConfig()
	}
}

// Load config file
func loadConfig(file string) {
	// Check if the config file exists
	if _, err := os.Stat(file); err != nil {
		die(2, fmt.Sprintf("Configuration file, %s, not found", file))
	}
	// Load the config file (you can implement this based on your file format)
}

// Show menu for configuration
func showMenu() {
	// Implement the menu logic here if needed
	fmt.Println("Menu is not implemented yet.")
}

// Question and answer dialog
func questionAnswer() {
	// Implement question and answer logic here if needed
	fmt.Println("Question and Answer is not implemented yet.")
}

// Archive and upload files
func archiveAndUpload(tarFile string) error {
	// Command to create the archive
	cmd := exec.Command("tar", "czf", tarFile, "--newer", syncFile, ".")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Failed to create archive: %w", err)
	}

	// SSH command to upload
	remoteCommand := fmt.Sprintf("cd \"%s\" || exit; tar -xpzf -", dest)
	sshCmd := exec.Command("ssh", "-p", port, "-l", user, host, remoteCommand)
	sshCmd.Stdin, _ = os.Open(tarFile)
	if err := sshCmd.Run(); err != nil {
		return fmt.Errorf("Failed to upload: %w", err)
	}

	// Update sync file timestamp
	if err := os.Chtimes(syncFile, time.Now(), time.Now()); err != nil {
		return fmt.Errorf("Failed to update sync file timestamp: %w", err)
	}

	return nil
}

// Print error message and exit
func die(code int, msg string) {
	fmt.Fprintf(os.Stderr, "%s %s: %s\n", scriptName, version, msg)
	os.Exit(code)
}

// Print configuration
func printConfig() {
	fmt.Println("Configuration:")
	for varName, description := range varInfo {
		fmt.Printf("%-35s ## %s\n", varName+"=\""+getVarValue(varName)+"\"", description)
	}
}

// Get variable value (returning a string)
func getVarValue(varName string) string {
	switch varName {
	case "host":
		return host
	case "port":
		return port
	case "dest":
		return dest
	case "user":
		return user
	case "source":
		return source
	case "archiveDir":
		return archiveDir
	case "syncFile":
		return syncFile
	default:
		return ""
	}
}

