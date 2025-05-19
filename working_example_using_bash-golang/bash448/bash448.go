package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	dbFileName      = "Filerec.sha256"
	E_DIR_NOMATCH   = 50
	E_BAD_DBFILE    = 51
)

var directory string

// Calculate the SHA256 checksum of a file
func sha256sum(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Set up the database
func setUpDatabase() error {
	f, err := os.Create(dbFileName)
	if err != nil {
		return fmt.Errorf("error creating database file: %v", err)
	}
	defer f.Close()

	// Write directory name to the first line
	_, err = f.WriteString(directory + "\n")
	if err != nil {
		return fmt.Errorf("error writing to database file: %v", err)
	}

	// Write checksums and filenames to the database
	files, err := filepath.Glob(filepath.Join(directory, "*"))
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		sum, err := sha256sum(file)
		if err != nil {
			return err
		}
		_, err = f.WriteString(sum + " " + filepath.Base(file) + "\n")
		if err != nil {
			return fmt.Errorf("error writing to database: %v", err)
		}
	}
	return nil
}

// Check the database
func checkDatabase() error {
	file, err := os.Open(dbFileName)
	if err != nil {
		fmt.Println("Unable to read checksum database file!")
		os.Exit(E_BAD_DBFILE)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0
	var directoryChecked string
	for scanner.Scan() {
		rec := scanner.Text()
		if lineNum == 0 {
			// First line contains the directory
			directoryChecked = rec
			if directoryChecked != directory {
				fmt.Println("Directories do not match up!")
				os.Exit(E_DIR_NOMATCH)
			}
		} else {
			// Following lines contain checksums and filenames
			fields := strings.Fields(rec)
			if len(fields) != 2 {
				continue
			}

			expectedChecksum := fields[0]
			fileName := fields[1]
			fullPath := filepath.Join(directory, fileName)
			actualChecksum, err := sha256sum(fullPath)
			if err != nil {
				fmt.Printf("Error calculating checksum for file %s: %v\n", fileName, err)
				continue
			}

			if expectedChecksum == actualChecksum {
				fmt.Printf("%s unchanged.\n", fileName)
			} else {
				fmt.Printf("%s : CHECKSUM ERROR!\n", fileName)
			}
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading database: %v", err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		directory, _ = os.Getwd() // If not specified, use the current directory
	} else {
		directory = os.Args[1]
	}

	clearCmd := exec.Command("clear")
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()

	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		fmt.Printf("Setting up database file, \"%s/%s\".\n\n", directory, dbFileName)
		if err := setUpDatabase(); err != nil {
			fmt.Printf("Error setting up database: %v\n", err)
			os.Exit(1)
		}
	}

	if err := checkDatabase(); err != nil {
		fmt.Printf("Error checking database: %v\n", err)
		os.Exit(1)
	}
}

