package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func countFilesWithOsWalk(path string) int {
	total := 0
	filepath.Walk(path, func(base string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			total++
		}
		return nil
	})
	return total
}

func countFilesWithScandir(path string) int {
	total := 0
	dir, _ := os.ReadDir(path)
	for _, entry := range dir {
		if entry.IsDir() {
			total += countFilesWithScandir(filepath.Join(path, entry.Name()))
		} else {
			total++
		}
	}
	return total
}

func countFilesWithPathlib(path string) int {
	total := 0
	dir, _ := os.ReadDir(path)
	for _, entry := range dir {
		if entry.IsDir() {
			total += countFilesWithPathlib(filepath.Join(path, entry.Name()))
		} else {
			total++
		}
	}
	return total
}

func main() {
	fmt.Println(countFilesWithOsWalk("."))
	fmt.Println(countFilesWithScandir("."))
	fmt.Println(countFilesWithPathlib("."))
}

// 1
// 1
// 1
