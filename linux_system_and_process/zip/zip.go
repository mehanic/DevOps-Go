package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// backupToZip backs up the entire contents of the specified folder into a zip file.
func backupToZip(folder string) error {
	// Get absolute path of the folder
	folder, err := filepath.Abs(folder)
	if err != nil {
		return err
	}

	// Generate the zip filename
	number := 1
	var zipFilename string
	for {
		zipFilename = fmt.Sprintf("%s_%d.zip", filepath.Base(folder), number)
		if _, err := os.Stat(zipFilename); os.IsNotExist(err) {
			break
		}
		number++
	}

	// Create the zip file
	fmt.Printf("Creating %s...\n", zipFilename)
	zipFile, err := os.Create(zipFilename)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the folder and add files to the zip archive
	err = filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create the zip header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name, err = filepath.Rel(filepath.Dir(folder), path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		// Create the writer for the file in the zip archive
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If the file is not a directory, write its content to the zip archive
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("Done.")
	return nil
}

func main() {
	err := backupToZip("archive-zip")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
