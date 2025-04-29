package main

import (
	"os"
	"path/filepath"
)

func deleteFileOs() {
	os.Remove("screenshots/screenshot1.png")
	// os.Remove("screenshots/screenshot1.png")
}

func deleteFilePathlib() {
	file := filepath.Clean("screenshots/screenshot2.png")
	os.Remove(file)
}

func deleteDirectoryOs() {
	os.Remove("screenshots/")
}

func deleteDirectoryPathlib() {
	directory := filepath.Clean("other-screenshots/")
	os.Remove(directory)
}

func deleteDirectoryIncludingSubdir() {
	os.RemoveAll("old-images/")
}

func main() {
	// deleteFileOs()
	// deleteFilePathlib()
	// deleteDirectoryOs()
	// deleteDirectoryPathlib()
	deleteDirectoryIncludingSubdir()
}
