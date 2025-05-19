package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
//	"strconv"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define the image types
	types := []string{"jpg", "png"}

	// Create the data directory if it doesn't exist
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// Generate 100 image files
	for i := 1; i <= 100; i++ {
		prefix := fmt.Sprintf("%03d", i)
		suffix := types[rand.Intn(len(types))]
		filename := filepath.Join("data", fmt.Sprintf("%s.%s", prefix, suffix))
		
		// Create the file
		_, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
	}

	// Rename the generated files
	count := 1
	err = filepath.Walk("data", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			extension := filepath.Ext(path)
			newName := fmt.Sprintf("./data/image-%03d%s", count, extension)
			err = os.Rename(path, newName)
			if err != nil {
				fmt.Println("Error renaming file:", err)
				return err
			}
			count++
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through data directory:", err)
	}
}

