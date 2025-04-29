package main

import (
	"io"
	"os"
)

func copyFile() {
	src, _ := os.Open("monster01.png")
	defer src.Close()

	dst, _ := os.Create("images/png")
	defer dst.Close()

	io.Copy(dst, src)
}

func copyFileWithMetadata() {
	src, _ := os.Open("monster02.png")
	defer src.Close()

	dst, _ := os.Create("images/png")
	defer dst.Close()

	io.Copy(dst, src)
}

func copyDirectory() {
	src := "images/"
	dst := "back-up-images/"
	os.MkdirAll(dst, os.ModePerm)

	files, _ := os.ReadDir(src)
	for _, file := range files {
		srcFile, _ := os.Open(src + file.Name())
		defer srcFile.Close()

		dstFile, _ := os.Create(dst + file.Name())
		defer dstFile.Close()

		io.Copy(dstFile, srcFile)
	}
}

func main() {
	// copyFile()
	// copyFileWithMetadata()
	copyDirectory()
}
