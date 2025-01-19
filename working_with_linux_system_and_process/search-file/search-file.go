package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	var reqFile string
	fmt.Print("Enter your file name to search: ")
	fmt.Scanln(&reqFile)

	if runtime.GOOS == "windows" {
		pdNames := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		vdNames := []string{}
		for _, eachDrive := range pdNames {
			if _, err := os.Stat(string(eachDrive) + ":\\"); err == nil {
				vdNames = append(vdNames, string(eachDrive)+":\\")
			}
		}
		fmt.Println(vdNames)
		for _, eachDrive := range vdNames {
			filepath.Walk(eachDrive, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && info.Name() == reqFile {
					fmt.Println(path)
				}
				return nil
			})
		}
	} else {
		filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && info.Name() == reqFile {
				fmt.Println(path)
			}
			return nil
		})
	}
}
