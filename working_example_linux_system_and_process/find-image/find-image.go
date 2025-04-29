package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func displayPngs() {
	pngFiles, _ := filepath.Glob("*.png")
	fmt.Println(pngFiles)
}

func findMonsterOne() {
	filteredItems, _ := filepath.Glob("*monster01*")
	fmt.Println(filteredItems)
}

func findMonsterOneInSubdirs() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Base(path) == "monster01" {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	displayPngs()
	findMonsterOne()
	findMonsterOneInSubdirs()
}
