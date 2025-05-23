package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// generateSHA1 generates SHA1 checksum for a single file.
func generateSHA1(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	hash := sha1.New()
	hash.Write(data)
	return fmt.Sprintf("%x  %s", hash.Sum(nil), filepath.Base(filePath)), nil
}

// verifyChecksum verifies the checksum of a file against a given hash.
func verifyChecksum(filePath, checksumFile string) (bool, error) {
	file, err := os.Open(checksumFile)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		hash, name := parts[0], parts[1]
		if name == filepath.Base(filePath) {
			checksum, err := generateSHA1(filePath)
			if err != nil {
				return false, err
			}
			if hash == strings.Split(checksum, " ")[0] {
				return true, nil
			}
		}
	}

	return false, scanner.Err()
}

// writeOutput appends the content to the output file.
func writeOutput(outputFile string, content string) error {
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content + "\n")
	return err
}

func main() {
	outputFile := "sha1sum.output"
	
	// Step 1: Generate single SHA1
	singleFile := "sha1.sh"
	checksumFile := "sha1.sha1"

	singleChecksum, err := generateSHA1(singleFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating checksum for %s: %v\n", singleFile, err)
		return
	}

	// Write to single checksum file
	err = ioutil.WriteFile(checksumFile, []byte(singleChecksum), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing checksum file: %v\n", err)
		return
	}

	// Output structure of single checksum file
	content, err := ioutil.ReadFile(checksumFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading checksum file: %v\n", err)
		return
	}
	
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if line != "" {
			writeOutput(outputFile, fmt.Sprintf("%d %s", i+1, line))
		}
	}

	// Step 2: Verify single SHA1
	ok, err := verifyChecksum(singleFile, checksumFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying checksum for %s: %v\n", singleFile, err)
		return
	}
	if ok {
		writeOutput(outputFile, fmt.Sprintf("%s checksum verified.", singleFile))
	} else {
		writeOutput(outputFile, fmt.Sprintf("%s checksum verification failed.", singleFile))
	}

	// Step 3: Generate SHA1 for all .txt files
	txtFiles, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding .txt files: %v\n", err)
		return
	}

	multiChecksumFile := "multi.sha1sum"
	for _, txtFile := range txtFiles {
		checksum, err := generateSHA1(txtFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating checksum for %s: %v\n", txtFile, err)
			return
		}
		// Write to multi checksum file
		err = writeOutput(multiChecksumFile, checksum)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing multi checksum file: %v\n", err)
			return
		}
	}

	// Output structure of multi checksum file
	multiContent, err := ioutil.ReadFile(multiChecksumFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading multi checksum file: %v\n", err)
		return
	}

	multiLines := strings.Split(string(multiContent), "\n")
	for i, line := range multiLines {
		if line != "" {
			writeOutput(outputFile, fmt.Sprintf("%d %s", i+1, line))
		}
	}

	// Step 4: Verify multi SHA1
	for _, txtFile := range txtFiles {
		ok, err := verifyChecksum(txtFile, multiChecksumFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error verifying checksum for %s: %v\n", txtFile, err)
			return
		}
		if ok {
			writeOutput(outputFile, fmt.Sprintf("%s checksum verified.", txtFile))
		} else {
			writeOutput(outputFile, fmt.Sprintf("%s checksum verification failed.", txtFile))
		}
	}

	// Final output
	finalOutput, err := ioutil.ReadFile(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading output file: %v\n", err)
		return
	}
	fmt.Println(string(finalOutput))
}

