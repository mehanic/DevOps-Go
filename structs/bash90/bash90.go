package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type FileInfo struct {
	Path string
	Size int64
	Hash string
}

// Function to calculate the MD5 checksum of a file
func getMD5Hash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func main() {
	dataDir := "./data/"
	files := make(map[int64][]FileInfo)

	// Read all files in the directory
	err := filepath.Walk(dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files[info.Size()] = append(files[info.Size()], FileInfo{Path: path, Size: info.Size()})
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	duplicateFiles := make(map[string][]string)

	// Identify duplicates based on size and checksum
	for _, fileInfos := range files {
		if len(fileInfos) < 2 {
			continue // Skip if there are no duplicates
		}

		// Check each file in the same size group
		for i := 0; i < len(fileInfos); i++ {
			hash, err := getMD5Hash(fileInfos[i].Path)
			if err != nil {
				fmt.Println("Error calculating hash for file:", fileInfos[i].Path)
				continue
			}

			// Group files by their hashes
			duplicateFiles[hash] = append(duplicateFiles[hash], fileInfos[i].Path)
		}
	}

	// Prepare to write to duplicate_files
	var duplicates []string
	for _, paths := range duplicateFiles {
		if len(paths) > 1 {
			duplicates = append(duplicates, paths...)
		}
	}

	// Sort and write unique duplicate file paths
	sort.Strings(duplicates)
	duplicateFilePath := "duplicate_files"
	ioutil.WriteFile(duplicateFilePath, []byte(fmt.Sprintf("%s\n", duplicates)), 0644)

	// Output duplicate files and remove them
	for _, filePath := range duplicates {
		fmt.Println("Removing duplicate:", filePath)
		err := os.Remove(filePath)
		if err != nil {
			fmt.Println("Error removing file:", filePath, err)
		}
	}
}

