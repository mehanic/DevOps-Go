package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	fmt.Print("Enter your directory path: ")
	fmt.Scanln(&path)

	if _, err := os.Stat(path); err == nil {
		listOfFilesDir, _ := os.ReadDir(path)
		fmt.Println("all files and dirs: ", listOfFilesDir)

		for _, fileOrDir := range listOfFilesDir {
			fdPath := path + string(os.PathSeparator) + fileOrDir.Name()
			if fileOrDir.IsDir() {
				fmt.Printf("%s is a directory\n", fdPath)
			} else {
				fmt.Printf("%s is a file\n", fdPath)
			}
		}
	} else {
		fmt.Println("please provide valid path")
		os.Exit(1)
	}
}
