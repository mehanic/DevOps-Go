package main

import (
	"fmt"
	"os"
)

func main() {
	var req_path string
	fmt.Print("Enter your directory path: ")
	fmt.Scanln(&req_path)

	if fileInfo, err := os.Stat(req_path); err == nil && fileInfo.IsDir() {
		all_f_ds, _ := os.ReadDir(req_path)
		if len(all_f_ds) == 0 {
			fmt.Printf("The given path is %s an empty path\n", req_path)
		} else {
			var req_ex string
			fmt.Print("Enter the required files extension .py/.sh/.log/.txt: ")
			fmt.Scanln(&req_ex)
			req_files := []string{}
			for _, each_f := range all_f_ds {
				if each_f.IsDir() {
					continue
				}
				if len(each_f.Name()) >= len(req_ex) && each_f.Name()[len(each_f.Name())-len(req_ex):] == req_ex {
					req_files = append(req_files, each_f.Name())
				}
			}
			if len(req_files) == 0 {
				fmt.Printf("There are no %s files in the location of %s\n", req_ex, req_path)
			} else {
				fmt.Printf("There are %d files in the location of %s with an extension of %s\n", len(req_files), req_path, req_ex)
				fmt.Printf("So, the files are: %v\n", req_files)
			}
		}
	} else {
		fmt.Printf("The given path %s is a file. Please pass only directory path\n", req_path)
	}
}
