package main

import (
	"fmt"
	"os"
)

func main() {
	var path string
	fmt.Print("Enter your path: ")
	fmt.Scanln(&path)

	if _, err := os.Stat(path); err == nil {
		fmt.Printf("Given path: %s is a valid path\n", path)
		if fileInfo, err := os.Stat(path); err == nil && !fileInfo.IsDir() {
			fmt.Println("and it is a file path")
		} else {
			fmt.Println("and it is a directory path")
		}
	} else {
		fmt.Printf("Given path: %s is not existing on this host\n", path)
	}
}
