package main

import (
	"archive/tar"
	"io"
	"os"
)

func createTarArchive() {
	files := []string{"monsters.csv", "Monster_contact_sheet.pdf"}
	tarFile, _ := os.Create("archive.tar")
	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	for _, file := range files {
		fileToTar, _ := os.Open(file)
		defer fileToTar.Close()
		stat, _ := fileToTar.Stat()
		header := &tar.Header{
			Name: file,
			Size: stat.Size(),
		}
		tarWriter.WriteHeader(header)
		io.Copy(tarWriter, fileToTar)
	}
}

func addToTarArchive() {
	tarFile, _ := os.OpenFile("archive.tar", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	fileToTar, _ := os.Open("monsters_card03.png")
	defer fileToTar.Close()
	stat, _ := fileToTar.Stat()
	header := &tar.Header{
		Name: "monsters_card03.png",
		Size: stat.Size(),
	}
	tarWriter.WriteHeader(header)
	io.Copy(tarWriter, fileToTar)
}

func extractTar() {
	tarFile, _ := os.Open("archive.tar")
	tarReader := tar.NewReader(tarFile)
	for {
		header, err := tarReader.Next()
		if err != nil {
			break
		}
		if header.Name == "monsters_card03.png" {
			fileToWrite, _ := os.Create("monsters_card03.png")
			defer fileToWrite.Close()
			io.Copy(fileToWrite, tarReader)
			break
		}
	}
}

func extractAll() {
	tarFile, _ := os.Open("archive.tar")
	tarReader := tar.NewReader(tarFile)
	os.Mkdir("extracted_monster_files", os.ModeDir)
	for {
		header, err := tarReader.Next()
		if err != nil {
			break
		}
		fileToWrite, _ := os.Create("extracted_monster_files/" + header.Name)
		defer fileToWrite.Close()
		io.Copy(fileToWrite, tarReader)
	}
}

func main() {
	createTarArchive()
	addToTarArchive()
	// extractTar()
	extractAll()
}
