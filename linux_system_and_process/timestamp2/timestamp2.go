package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func formatDatetime(timestamp int64) string {
	utcTimestamp := time.Unix(timestamp, 0)
	formattedDate := utcTimestamp.Format("02 Jan 2006 15 04 05")
	return formattedDate
}

func displayEntriesInDirectory(directory string) {
	entries, _ := os.ReadDir(directory)
	for _, entry := range entries {
		fmt.Println("Name: ", entry.Name())
		info, _ := entry.Info()
		fmt.Println("Creation Time:", formatDatetime(info.Sys().(*syscall.Stat_t).Birthtimespec.Sec))
		// Windows: fmt.Println("Creation Time: ", info.Sys().(*syscall.Stat_t).Ctimespec.Sec)
		fmt.Println("Last Access Time:", formatDatetime(info.Sys().(*syscall.Stat_t).Atimespec.Sec))
		fmt.Println("Size:", info.Size())
	}
}

func displayDirectories(directory string) {
	entries, _ := os.ReadDir(directory)
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("Directory Name:", entry.Name())
		}
	}
}

func displayFiles(directory string) {
	entries, _ := os.ReadDir(directory)
	for _, entry := range entries {
		if !entry.IsDir() {
			fmt.Println("File Name:", entry.Name())
		}
	}
}

func main() {
	// displayEntriesInDirectory("artwork/")
	displayDirectories("artwork/")
	displayFiles("artwork/")
}
