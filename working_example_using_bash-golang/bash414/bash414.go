package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var (
	nFiles  int
	fileType string
	directory string
	name     string
	ext      string
	unit     string
	tmpFile  string
	lower    int
	upper    int
)

func init() {
	// Initialize defaults
	nFiles = 3
	fileType = "binary"
	directory = "qa-data"
	name = "garbage"
	ext = ".bin"
	unit = "M"
	tmpFile = "/tmp/tmp.datamaker.sh"

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
}

// Get a random number in a range
func getRandomNumber(step, variance, upper int) int {
	return rand.Intn((upper-variance)/step+1)*step + variance
}

// Create binary files with random sizes
func createBinaryFiles(numFiles, lower, upper int) {
	ext = ".bin"
	for counter := 0; counter < numFiles; counter++ {
		rSize := getRandomNumber(1, lower, upper)
		fmt.Println("Creating binary file... please wait...")
		filePath := filepath.Join(directory, fmt.Sprintf("%s%d%s", name, counter, ext))
		cmd := exec.Command("dd", "if=/dev/zero", fmt.Sprintf("of=%s", filePath), fmt.Sprintf("bs=%d%s", rSize, unit), "count=1")
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error creating binary file:", err)
		}
	}
}

// Create text files with random ranges of data
func createTextFiles(numFiles, variance, upper int) {
	ext = ".txt"
	middle := upper / 4
	lower := middle - variance
	if lower < 0 {
		lower = -lower + 1
	}

	for counter := 0; counter < numFiles; counter++ {
		var end, start int
		validRange := false

		for !validRange {
			end = getRandomNumber(1, variance, upper)
			start = getRandomNumber(1, variance, lower)

			if end > start {
				validRange = true
			} else {
				fmt.Println("Generating random values... please wait...")
			}
		}

		fmt.Println("Creating text file... please wait...")
		filePath := filepath.Join(directory, fmt.Sprintf("%s%d%s", name, counter, ext))
		cmd := exec.Command("dd", fmt.Sprintf("if=%s", tmpFile), fmt.Sprintf("of=%s", filePath), fmt.Sprintf("seek=%d", start), "bs=1", fmt.Sprintf("count=%d", end-start))
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error creating text file:", err)
		}
	}
}

func main() {
	// Parse command-line arguments
	flag.StringVar(&fileType, "t", "binary", "File type: binary or text")
	flag.IntVar(&nFiles, "n", 3, "Number of files to create")
	flag.IntVar(&lower, "l", 1, "Lower bound of file size")
	flag.IntVar(&upper, "u", 10, "Upper bound of file size")
	flag.Parse()

	// Validation
	if lower <= 0 || upper <= 0 || nFiles <= 0 {
		fmt.Println("ERROR: Invalid arguments. Ensure -n, -l, and -u are greater than 0.")
		os.Exit(1)
	}

	// Create the target directory if it doesn't exist
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		os.Mkdir(directory, 0755)
	}

	// Create files based on type
	switch fileType {
	case "binary":
		createBinaryFiles(nFiles, lower, upper)
	case "text":
		createTextFiles(nFiles, lower, upper)
	default:
		fmt.Println("ERROR: Unknown file type. Use 'binary' or 'text'.")
		os.Exit(1)
	}

	fmt.Println("DONE!")
}

