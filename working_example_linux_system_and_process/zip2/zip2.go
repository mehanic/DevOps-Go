package main

import (
	"archive/zip"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	zipFile := "example.zip"
	dir := strings.TrimSuffix(zipFile, filepath.Ext(zipFile))
	os.Chdir("automate_online-materials")

	exampleZip, err := zip.OpenReader(zipFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer exampleZip.Close()

	var spamFile *zip.File
	for _, f := range exampleZip.File {
		if f.Name == "spam.txt" {
			spamFile = f
			break
		}
	}

	if spamFile != nil {
		fmt.Printf("Compressed file is %fx smaller!\n", round(float64(spamFile.UncompressedSize64)/float64(spamFile.CompressedSize64), 2))
	}

	os.MkdirAll(dir, os.ModePerm)
	extractAll(exampleZip, dir)
	extractFile(spamFile, dir+"_r2")
}

func round(val float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	return math.Round(val*pow) / pow
}

func extractAll(zipReader *zip.ReadCloser, dest string) {
	for _, f := range zipReader.File {
		extractFile(f, dest)
	}
}

func extractFile(f *zip.File, dest string) {
	rc, err := f.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rc.Close()

	fpath := filepath.Join(dest, f.Name)
	if f.FileInfo().IsDir() {
		os.MkdirAll(fpath, os.ModePerm)
	} else {
		os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
		outFile, err := os.Create(fpath)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer outFile.Close()
		_, err = io.Copy(outFile, rc)
		if err != nil {
			fmt.Println(err)
		}
	}
}
