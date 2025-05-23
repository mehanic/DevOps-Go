package main

import (
	"fmt"
	"os"
//	"path/filepath"
)

// changedir changes the current working directory and checks various permissions
func changedir(dirName string) error {
	// Check if the directory exists
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		return fmt.Errorf("Dir: %s does not exist", dirName)
	}

	// Check if the directory is readable
	fileInfo, err := os.Stat(dirName)
	if err != nil {
		return fmt.Errorf("error checking directory: %v", err)
	}

	// Check if we have execute permissions on the directory
	if fileInfo.Mode()&os.ModePerm&0111 == 0 {
		return fmt.Errorf("Dir: cannot cd to %s", dirName)
	}

	// Check if the directory is writable
	if fileInfo.Mode()&os.ModePerm&0200 == 0 {
		return fmt.Errorf("Dir: %s not writable", dirName)
	}

	// Change to the specified directory
	if err := os.Chdir(dirName); err != nil {
		return fmt.Errorf("error changing directory: %v", err)
	}

	fmt.Printf("Present directory %s\n", dirName)
	return nil
}

func main() {
	// Example usage
	dirName := "/data/app/" // Change this to the directory you want to check

	if err := changedir(dirName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

