package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	folderPath := ""
	destinationPath := ""

	// Read the directories in the folderPath
	folders, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, folder := range folders {
		if folder.IsDir() {
			// Read files in the subdirectory
			files, err := ioutil.ReadDir(filepath.Join(folderPath, folder.Name()))
			if err != nil {
				fmt.Println("Error reading subdirectory:", err)
				continue
			}

			for _, file := range files {
				srcPath := filepath.Join(folderPath, folder.Name(), file.Name())
				dstPath := filepath.Join(destinationPath, file.Name())

				// Open the image file
				srcFile, err := os.Open(srcPath)
				if err != nil {
					fmt.Println("Error opening file:", err)
					continue
				}
				defer srcFile.Close()

				// Decode the image
				var img image.Image
				switch strings.ToLower(filepath.Ext(file.Name())) {
				case ".jpg", ".jpeg":
					img, _, err = image.Decode(srcFile)
				case ".png":
					img, _, err = image.Decode(srcFile)
				default:
					fmt.Println("Unsupported file format:", file.Name())
					continue
				}
				if err != nil {
					fmt.Println("Error decoding image:", err)
					continue
				}

				// Create the destination file
				dstFile, err := os.Create(dstPath)
				if err != nil {
					fmt.Println("Error creating file:", err)
					continue
				}
				defer dstFile.Close()

				// Encode and save the image
				switch strings.ToLower(filepath.Ext(file.Name())) {
				case ".jpg", ".jpeg":
					err = jpeg.Encode(dstFile, img, nil)
				case ".png":
					err = png.Encode(dstFile, img)
				}
				if err != nil {
					fmt.Println("Error encoding image:", err)
					continue
				}
			}
		}
	}

	fmt.Println("All Files Moved.")
}

