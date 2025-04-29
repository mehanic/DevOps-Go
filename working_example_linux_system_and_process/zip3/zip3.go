package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func readZip() {
	archive, err := zip.OpenReader("archive.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	for _, file := range archive.File {
		log.Println(file.Name)
	}
}

func retrieveFileInfoZip() {
	archive, err := zip.OpenReader("archive.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	for _, file := range archive.File {
		if file.Name == "description01.txt" {
			log.Println(file.FileInfo())
		}
	}
}

func readFileInZip() {
	archive, err := zip.OpenReader("archive.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	for _, file := range archive.File {
		if file.Name == "description01.txt" {
			f, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			buf := make([]byte, file.FileInfo().Size())
			_, err = f.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(buf))
		}
	}
}

func extractFile(archive, fileToExtract string) {
	myArchive, err := zip.OpenReader(archive)
	if err != nil {
		log.Fatal(err)
	}
	defer myArchive.Close()

	for _, file := range myArchive.File {
		if file.Name == fileToExtract {
			f, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			outFile, err := os.Create(fileToExtract)
			if err != nil {
				log.Fatal(err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, f)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func extractAll() {
	archive, err := zip.OpenReader("archive.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer archive.Close()

	for _, file := range archive.File {
		f, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		outFile, err := os.Create("extracted_files/" + file.Name)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()

		_, err = io.Copy(outFile, f)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func main() {
	//readZip()
	//retrieveFileInfoZip()
	//readFileInZip()
	//extractFile("archive.zip", "description01.txt")
	extractAll()
}
