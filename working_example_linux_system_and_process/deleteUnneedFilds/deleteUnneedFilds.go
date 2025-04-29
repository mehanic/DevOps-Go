package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir := "/home"
	size := int64(1024) // size in MB

	multiplicator := int64(1000 * 1000)
	size = size * multiplicator

	largeFiles := []struct {
		path string
		size float64
	}{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore mounted remote storage.
		if filepath.HasPrefix(path, "/home/iliya/dune-mountpoint") {
			return filepath.SkipDir
		}

		if !info.IsDir() && info.Size() > size {
			largeFiles = append(largeFiles, struct {
				path string
				size float64
			}{
				path: path,
				size: float64(info.Size()) / float64(multiplicator),
			})
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", dir, err)
		return
	}

	for _, file := range largeFiles {
		fmt.Printf("%s %.6f MB\n", file.path, file.size)
	}

	
}


// send2trash moves the file at the given path to the system's trash or recycle bin.
// If the operation is not supported or fails, an error is returned.
func send2trash(path string) error {
	// Implementation goes here
	return nil
}
