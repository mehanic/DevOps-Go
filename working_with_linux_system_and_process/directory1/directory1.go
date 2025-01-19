package main

import (
	"fmt"
	"os"
	"path/filepath"
)
// filepath.Walk: Walks the directory tree rooted at "artwork/".
// Callback Function: Called for each file or directory in the tree. 
//It prints the directory or file information.
// os.ReadDir: Reads the directory named by dirpath and returns a list of 
//directory entries sorted by filename.

func topDownWalk() {
	err := filepath.Walk("artwork/", func(dirpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Println("Directory:", dirpath)
			fmt.Println("Includes these directories")
			dirnames, _ := os.ReadDir(dirpath)
			for _, dir := range dirnames {
				fmt.Println(dir.Name())
			}
		} else {
			fmt.Println("Includes these files")
			fmt.Println(info.Name())
		}
		fmt.Println()
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
//This function is identical to topDownWalk in the provided code. However, the name suggests 
//it should walk the directory tree in a bottom-up manner. Currently, both functions have the same 
//implementation and behavior.
func bottomUpWalk() {
	err := filepath.Walk("artwork/", func(dirpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Println("Directory:", dirpath)
			fmt.Println("Includes these directories")
			dirnames, _ := os.ReadDir(dirpath)
			for _, dir := range dirnames {
				fmt.Println(dir.Name())
			}
		} else {
			fmt.Println("Includes these files")
			fmt.Println(info.Name())
		}
		fmt.Println()
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// topDownWalk()
	bottomUpWalk()
}
