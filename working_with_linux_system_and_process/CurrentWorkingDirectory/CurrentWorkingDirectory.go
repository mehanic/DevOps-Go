package main

import (
	"fmt"
	"os"
)

func displayCwd() {
	cwd, _ := os.Getwd()
	fmt.Println("Current Working Directory: ", cwd)
}

func upOneDirectoryLevel() {
	_ = os.Chdir("../")
}

func displayEntriesInDirectory(directory string) {
	entries, _ := os.ReadDir(directory)
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
}

func main() {
	displayCwd()
	upOneDirectoryLevel()
	displayCwd()
	displayEntriesInDirectory("images/")
}
