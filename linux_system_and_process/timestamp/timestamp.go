package main

import (
	"fmt"
	"os"
	"time"
)

func formatDatetime(timestamp int64) string {
	utcTimestamp := time.Unix(timestamp, 0)
	formattedDate := utcTimestamp.Format("02 Jan 2006 15 04 05")
	return formattedDate
}

func displayEntriesInDirectory(directory string) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, entry := range entries {
		fmt.Println("Name:", entry.Name())
	}
}

func main() {
	fmt.Println(formatDatetime(time.Now().Unix()))
	displayEntriesInDirectory(".")
}
