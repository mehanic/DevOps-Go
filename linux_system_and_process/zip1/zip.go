package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func createZip(files []string) {
	newZipFile, _ := os.Create("archive.zip")
	defer newZipFile.Close()

	archive := zip.NewWriter(newZipFile)
	defer archive.Close()

	for _, file := range files {
		archiveFile, _ := os.Open(file)
		defer archiveFile.Close()

		info, _ := archiveFile.Stat()
		header, _ := zip.FileInfoHeader(info)
		header.Name = filepath.Base(file)

		writer, _ := archive.CreateHeader(header)
		io.Copy(writer, archiveFile)
	}
}
func createZipNestedDir(directoryPath string) {
	archive, err := os.Create("directory_archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	err = filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(path)

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	if err != nil {
		panic(err)
	}
}
func main() {
	//files := []string{"png/monster01.png", "png/monster02.png", "png/monster03.png"}
	//createZip(files)
	createZipNestedDir("svg")
}
