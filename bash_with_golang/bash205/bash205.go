package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	// Create a regex that matches files with the American date format.
	datePattern := regexp.MustCompile(`^(.*?) ((0|1)?\d)-((0|1|2|3)?\d)-((19|20)\d\d)(.*?)$`)

	// Get the current working directory.
	absWorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Loop over the files in the working directory.
	files, err := ioutil.ReadDir(absWorkingDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		amerFilename := file.Name()
		mo := datePattern.FindStringSubmatch(amerFilename)

		// Skip files without a date.
		if mo == nil {
			continue
		}

		// Get the different parts of the filename.
		beforePart := mo[1]
		monthPart := mo[2]
		dayPart := mo[4]
		yearPart := mo[6]
		afterPart := mo[8]

		// Form the European-style filename.
		euroFilename := fmt.Sprintf("%s%s-%s-%s%s", beforePart, dayPart, monthPart, yearPart, afterPart)

		// Get the full, absolute file paths.
		amerFullPath := filepath.Join(absWorkingDir, amerFilename)
		euroFullPath := filepath.Join(absWorkingDir, euroFilename)

		// Print the renaming action.
		fmt.Printf("Renaming \"%s\" to \"%s\"...\n", amerFullPath, euroFullPath)

		// Actually rename the files (uncomment this after testing).
		err := os.Rename(amerFullPath, euroFullPath)
		if err != nil {
			fmt.Println("Error renaming file:", err)
		}
	}
}

