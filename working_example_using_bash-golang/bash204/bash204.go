package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

const squareFitSize = 300
const logoFilename = "catlogo.png"

func main() {
	// Open the logo image
	logoIm, err := imaging.Open(logoFilename)
	if err != nil {
		log.Fatalf("Failed to open logo image: %v", err)
	}
	logoWidth := logoIm.Bounds().Dx()
	logoHeight := logoIm.Bounds().Dy()

	// Create output directory if it doesn't exist
	err = os.MkdirAll("withLogo", os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Loop through all the files in the current directory
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-image files
		if info.IsDir() || (!filepath.Ext(path).MatchString(".png") && !filepath.Ext(path).MatchString(".jpg")) || info.Name() == logoFilename {
			return nil
		}

		// Open the image file
		im, err := imaging.Open(path)
		if err != nil {
			fmt.Printf("Failed to open image %s: %v\n", path, err)
			return nil
		}

		width := im.Bounds().Dx()
		height := im.Bounds().Dy()

		// Resize the image if necessary
		if width > squareFitSize || height > squareFitSize {
			if width > height {
				height = (squareFitSize * height) / width
				width = squareFitSize
			} else {
				width = (squareFitSize * width) / height
				height = squareFitSize
			}

			// Resize image
			fmt.Printf("Resizing %s...\n", path)
			im = imaging.Resize(im, width, height, imaging.Lanczos)
		}

		// Add the logo to the lower-right corner
		fmt.Printf("Adding logo to %s...\n", path)
		im = imaging.Overlay(im, logoIm, image.Pt(width-logoWidth, height-logoHeight), 1.0)

		// Save the image to the output directory
		outputPath := filepath.Join("withLogo", info.Name())
		err = imaging.Save(im, outputPath)
		if err != nil {
			fmt.Printf("Failed to save image %s: %v\n", outputPath, err)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Failed to process files: %v", err)
	}
	fmt.Println("Done.")
}

