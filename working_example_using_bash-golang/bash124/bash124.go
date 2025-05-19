package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"strconv"
	"time"
)

func main() {
	// Set constants
	backupDir := fmt.Sprintf("/var/rhevm-backups/%s", getHostname())
	backupReports := false // Set to true if using RHEV reports, false otherwise
	daysToRetainBackups := 7

	// Set backup date format
	backupDate := time.Now().Format("Mon_Jan_2_15_04_05_MST_2006")

	// The backup files
	configBackupFile := filepath.Join(backupDir, fmt.Sprintf("config_%s.tar.gz", backupDate))
	reportsConfigBackupFile := filepath.Join(backupDir, fmt.Sprintf("reports_config_%s.tar.gz", backupDate))

	// Ensure the script is run as root
	if os.Geteuid() != 0 {
		fmt.Println("Please run with sudo or as root.")
		os.Exit(1)
	}

	// Create the backup directory if it doesn't exist
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		err := os.MkdirAll(backupDir, 0775)
		if err != nil {
			fmt.Println("Error creating backup directory:", err)
			os.Exit(1)
		}
		if err := os.Chmod(backupDir, 0775); err != nil {
			fmt.Println("Error setting permissions:", err)
		}
		if err := os.Chown(backupDir, -1, getGroupID("wheel")); err != nil {
			fmt.Println("Error changing group ownership:", err)
		}
	}

	// Backup the database
	runCommand("/usr/share/ovirt-engine/dbscripts/backup.sh", "-s", "localhost", "-d", "engine", "-u", "postgres", "-l", backupDir)

	// Compress the backup
	runCommand("gzip", filepath.Join(backupDir, "*.sql"))

	// Backup config files
	configFiles := []string{
		"/etc/ovirt-engine/", "/etc/sysconfig/ovirt-engine", "/etc/yum/pluginconf.d/versionlock.list",
		"/etc/pki/ovirt-engine/", "/usr/share/ovirt-engine/conf/iptables.*", "/usr/share/ovirt-engine/dbscripts/create_db.sh.log",
		"/var/lib/ovirt-engine/backups", "/var/lib/ovirt-engine/deployments", "/root/.rnd",
	}
	createTar(configBackupFile, configFiles)

	// Backup reports if needed
	if backupReports {
		configFilesForReports := []string{
			"/usr/share/ovirt-engine-reports/reports/users/rhevm-002dadmin.xml",
			"/usr/share/ovirt-engine-reports/default_master.properties",
			"/usr/share/jasperreports-server-pro/buildomatic",
		}
		createTar(reportsConfigBackupFile, configFilesForReports)
	}

	// Remove old backup files
	runCommand("tmpwatch", fmt.Sprintf("--mtime %dd", daysToRetainBackups), backupDir)
}

// Helper to run shell commands
func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running %s: %v\n", name, err)
		os.Exit(1)
	}
}

// Helper to create tar files
func createTar(backupFile string, files []string) {
	args := append([]string{"czPf", backupFile}, files...)
	runCommand("tar", args...)
}

// Helper to get hostname
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		os.Exit(1)
	}
	return hostname
}

// Helper to get group ID
func getGroupID(groupName string) int {
	group, err := exec.Command("getent", "group", groupName).Output()
	if err != nil {
		fmt.Println("Error getting group ID:", err)
		os.Exit(1)
	}

	// Parse group ID from output
	parts := strings.Split(string(group), ":")
	if len(parts) >= 3 {
		id := strings.TrimSpace(parts[2])
		gid, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println("Error parsing group ID:", err)
			os.Exit(1)
		}
		return gid
	}
	return -1
}

