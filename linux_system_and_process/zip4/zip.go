package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func main() {
	newZip, err := zip.OpenReader("result.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newZip.Close()

	for _, file := range newZip.File {
		fmt.Println(file.Name)
	}

	os.Exit(0)

	newZipWriter, err := os.Create("result.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newZipWriter.Close()

	zipWriter := zip.NewWriter(newZipWriter)
	defer zipWriter.Close()

	myPetsFile, err := os.Open("myPets.py")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer myPetsFile.Close()

	w, err := zipWriter.Create("myPets.py")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = io.Copy(w, myPetsFile)
	if err != nil {
		fmt.Println(err)
		return
	}
}
