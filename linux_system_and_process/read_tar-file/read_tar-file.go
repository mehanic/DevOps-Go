package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func readTarfile() {
	file, err := os.Open("archive.tar.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := tar.NewReader(file)
	for {
		header, err := reader.Next()
		if err != nil {
			break
		}
		fmt.Println(header.Name)
		fmt.Println(header.Size)
		fmt.Println(header.Mode)
	}
}

func readFileInTar() {
	file, err := os.Open("archive.tar.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := tar.NewReader(file)
	for {
		header, err := reader.Next()
		if err != nil {
			break
		}
		if header.Name == "descriptions/description01.txt" {
			data := make([]byte, header.Size)
			_, err := reader.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(data))
			break
		}
	}
}

func main() {
	//readTarfile()
	readFileInTar()
}
