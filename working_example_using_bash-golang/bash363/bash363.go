package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if 'work' is a directory
	if _, err := os.Stat("work"); os.IsNotExist(err) {
		fmt.Println("work is not a directory.")
	} else if stat, err := os.Stat("work"); err == nil && stat.IsDir() {
		fmt.Println("work is a directory.")
	} else {
		fmt.Println("work exists but is not a directory.")
	}

	// Check if 'test.txt' is a file
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("test.txt does not exist.")
	} else if stat, err := os.Stat("test.txt"); err == nil && stat.Mode().IsRegular() {
		fmt.Println("test.txt is a file.")
	} else {
		fmt.Println("test.txt exists but is not a file.")
	}

	// Check if 'test.txt' has size greater than 0
	if stat, err := os.Stat("test.txt"); err == nil && stat.Size() > 0 {
		fmt.Println("test.txt has size greater than 0.")
	} else {
		fmt.Println("test.txt is either empty or does not exist.")
	}
}

