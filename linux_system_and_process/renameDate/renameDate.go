package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// Function to rename files with American date format (MM-DD-YYYY) to European (DD-MM-YYYY)
func main() {
	// Regular expression to match the American date format in filenames
	datePattern := regexp.MustCompile(`^(.*?)((0|1)?\d)-((0|1|2|3)?\d)-((19|20)\d\d)(.*?)$`)

	// Get the absolute working directory
	absWorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	// Loop over the files in the current working directory
	err = filepath.Walk(absWorkingDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if the file name matches the American date format
		amerFilename := info.Name()
		matches := datePattern.FindStringSubmatch(amerFilename)
		if matches == nil {
			return nil // Skip files without a matching date format
		}

		// Extract parts of the filename
		beforePart := matches[1]
		monthPart := matches[2]
		dayPart := matches[4]
		yearPart := matches[6]
		afterPart := matches[8]

		// Form the European-style filename (DD-MM-YYYY)
		euroFilename := beforePart + dayPart + "-" + monthPart + "-" + yearPart + afterPart

		// Get the full absolute paths for the original and new filenames
		oldPath := filepath.Join(absWorkingDir, amerFilename)
		newPath := filepath.Join(absWorkingDir, euroFilename)

		// Print renaming operation for testing purposes
		fmt.Printf("Renaming \"%s\" to \"%s\"...\n", oldPath, newPath)

		// Uncomment the following line to actually rename the file after testing
		// err = os.Rename(oldPath, newPath)
		// if err != nil {
		//     fmt.Println("Error renaming file:", err)
		// }

		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directory:", err)
	}
}


//Running the Go Program:
//Place this Go script in the same directory as the files you want to rename.
//Run the Go program. It will print the intended renaming operations without actually renaming the files (to avoid accidental changes during testing).
//Once you are sure it works, uncomment the os.Rename() line to enable the actual renaming.
///This Go program will successfully rename files with American date formats (MM-DD-YYYY) to European date formats (DD-MM-YYYY).
