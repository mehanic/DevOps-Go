package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Slice to hold the unique file hashes
var fileArray []string

// Function to add file hash to the slice
func addFile(hash string) {
	fileArray = append(fileArray, hash)
}

// Function to delete a file
func delFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Printf("Error deleting file %s: %v\n", file, err)
	}
}

// Function to compute the SHA-512 hash of a file
func hashFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := sha512.New()
	_, err = io.Copy(hasher, f)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// Function to search for the file's hash in the array
func searchFile(file string) (string, int) {
	hash, err := hashFile(file)
	if err != nil {
		fmt.Printf("Error hashing file %s: %v\n", file, err)
		return "", 1
	}

	if len(fileArray) == 0 {
		// Edge case: First file, no files in array yet
		return hash, 0
	}

	for _, existingHash := range fileArray {
		if existingHash == hash {
			// Return 1 if the file is found in the array
			return "", 1
		}
	}

	// Return the hash if it's not found in the array
	return hash, 0
}

// Function to search directory and remove duplicates
func beginSearchAndDeduplication(dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}

		if !info.IsDir() {
			hash, ret := searchFile(path)
			if ret == 1 {
				// Duplicate file, delete it
				delFile(path)
			} else if hash != "" {
				// Unique file, add its hash
				addFile(hash)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the directory %s: %v\n", dir, err)
	}
}

// Function to display the file hashes in the array
func dumpArray() {
	for i, hash := range fileArray {
		fmt.Printf("#%d %s\n", i+1, hash)
	}
}

func main() {
	// Get directory input from the user
	var dir string
	fmt.Println("Enter directory name to begin searching and deduplicating:")
	fmt.Scanln(&dir)

	// Start searching and deduplicating
	beginSearchAndDeduplication(dir)

	// Dump the array of unique file hashes
	dumpArray()
}

